package main

import (
	"fmt"
	"os"
)

func writeAudioHijackFileOutput(track Track, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if track.Name != "" {
		fmt.Fprintln(f, "Title:", track.Name)
	}

	if track.Artist != "" {
		fmt.Fprintln(f, "Artist:", track.Artist)
	}

	if track.ImagePath != "" {
		fmt.Fprintln(f, "Artwork:", "file://"+track.ImagePath)
	}

}
