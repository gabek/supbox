package main

type RekordboxConfig struct {
	Options  [][]string `json:"options"`
	Defaults struct {
		Mode         string `json:"mode"`
		Connectivity struct {
			URL string `json:"url"`
		} `json:"connectivity"`
		ClockServer struct {
			Urls []string `json:"urls"`
		} `json:"clock_server"`
	} `json:"defaults"`
}