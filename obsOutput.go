package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func writeOBSFilesOutput(track Track, filenameTemplate string) {
	artistFilename := strings.Replace(filenameTemplate, "*", "artist.txt", 1)
	writeTextToFilePath(track.Artist, artistFilename)

	trackFilename := strings.Replace(filenameTemplate, "*", "track.txt", 1)
	writeTextToFilePath(track.Name, trackFilename)

	imageFilename := strings.Replace(filenameTemplate, "*", "track.jpg", 1)
	copyFileContents(track.ImagePath, imageFilename)
}

func writeTextToFilePath(text string, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(f, text)
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
