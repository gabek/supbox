package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func getDataPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to determine your home directory to location options.json")
	}

	return filepath.Join(homeDir, "/Library/Pioneer/rekordbox")
}

func getRekordboxConfig(path string) RekordboxConfig {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to determine your home directory to location options.json")
	}
	optionsFilePath := filepath.Join(homeDir, "/Library/Application Support/Pioneer/rekordboxAgent/storage/", "options.json")

	// read file
	data, err := ioutil.ReadFile(optionsFilePath)
	if err != nil {
		fmt.Print(err)
	}

	// json data
	var obj RekordboxConfig

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getEncryptedPassword(asarPath string) string {
	f, err := os.Open(asarPath)
	if err != nil {
		panic("Cnnot open asar file: " + asarPath)
	}

	// Search for the password
	data, err := ioutil.ReadAll(f)
	re := regexp.MustCompile(`pass: \".*\"`)
	result := re.FindAllString(string(data), 10)[0]

	password := strings.Split(result, ": ")[1]

	// Remove the quotation marks
	password = strings.Replace(password, `"`, "", -1)

	return password

}

func getAsarFilePath(root string) string {
	return filepath.Join(root, "/Contents/MacOS/rekordboxAgent.app/Contents/Resources/app.asar")
}

func getEncryptedPasswordDataFromConfig(config RekordboxConfig) string {
	if len(config.Options) < 2 {
		panic("Unable to read Rekordbox Config file to get password.")
	}

	if len(config.Options[1]) < 2 {
		panic("Unable to read Rekordbox Config file to get password.")
	}

	return config.Options[1][1]
}

func getDatabaseFilePath(root string) string {
	return filepath.Join(root, "master.db")
}

func getImagePath(dataPath string, imagePath string) string {
	return filepath.Join(dataPath, "share", imagePath)
}
