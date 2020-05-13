package main

import (
	"database/sql"
	"fmt"
)

func getDatabaseDSN(filePath string, encryptionKey string) string {
	dsn := fmt.Sprintf("file:"+filePath+"?_key='%s'", encryptionKey)
	return dsn
}

func getRecentTrack(db *sql.DB, config Config) Track {
	row := db.QueryRow("SELECT h.ID, Title, Name, ImagePath FROM djmdSongHistory AS h JOIN djmdContent AS c on h.ContentID = c.ID LEFT JOIN djmdArtist as a on c.ArtistID = a.ID GROUP BY h.created_at ORDER BY h.created_at DESC LIMIT 1")

	var ID string
	var Title string
	var Name string
	var ImagePath string

	row.Scan(&ID, &Title, &Name, &ImagePath)

	if ImagePath != "" {
		dataPath := getDataPath()
		ImagePath = getImagePath(dataPath, ImagePath)
	}

	return Track{ID: ID, Artist: Name, Name: Title, ImagePath: ImagePath}
}
