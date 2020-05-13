package main

type Track struct {
	ID        string `json:-`
	Artist    string `json:"artist"`
	Name      string `json:"track"`
	ImagePath string `json:"imagePath"`
}
