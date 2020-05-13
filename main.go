package main

import (
	"encoding/base64"
	"log"

	"database/sql"
	"time"
)

var config Config = getConfig()

var database *sql.DB

var currentTrackID string

func main() {
	// Files and paths
	rekordboxConfig := getRekordboxConfig(config.Rekordbox.OptionsFile)
	asarFilePath := getAsarFilePath(config.Rekordbox.ApplicationPath)
	dataPath := getDataPath()
	databaseFilePath := getDatabaseFilePath(dataPath)

	// Database decryption
	encodedPasswordData := getEncryptedPasswordDataFromConfig(rekordboxConfig)
	decodedPasswordData, err := base64.StdEncoding.DecodeString(encodedPasswordData)
	passwordString := getEncryptedPassword(asarFilePath)
	password := []byte(passwordString)
	decryptedBytes := decrypt(decodedPasswordData, password)
	encryptionKey := string(decryptedBytes)

	// Open the Database
	dsn := getDatabaseDSN(databaseFilePath, encryptionKey)
	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		panic(err)
	}

	database = db
	defer database.Close()

	// Start polling
	pollingInterval, err := time.ParseDuration(config.PollingInterval)

	if err != nil {
		panic(err)
	}

	startTimer(pollingInterval)
}

func startTimer(pollingInterval time.Duration) {
	run()

	tick := time.Tick(pollingInterval)
	for range tick {
		run()
	}
}

func run() {
	track := getRecentTrack(database, config)
	if track.ID == currentTrackID {
		return
	}

	currentTrackID = track.ID

	log.Printf("%+v\n", track)

	if config.Output.AudioHijackStyleFile != "" {
		writeAudioHijackFileOutput(track, config.Output.AudioHijackStyleFile)
	}
	if config.Output.JSONStyleFile != "" {
		writeNowPlayingJSONOutput(track, config.Output.JSONStyleFile)
	}

	if config.Output.OBSStyleFileTemplate != "" {
		writeOBSFilesOutput(track, config.Output.OBSStyleFileTemplate)
	}
}
