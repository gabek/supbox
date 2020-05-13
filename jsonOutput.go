package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getJSONForTrack(track Track) string {
	b, err := json.Marshal(track)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func writeNowPlayingJSONOutput(track Track, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	json := getJSONForTrack(track)
	fmt.Fprintln(f, json)
}
