package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Rekordbox       Rekordbox `yaml:"rekordbox"`
	Output          Output    `yaml:"output"`
	PollingInterval string    `yaml:"pollingInterval"`
}

type Rekordbox struct {
	OptionsFile     string `yaml:"optionsFile"`
	ApplicationPath string `yaml:"application"`
}

type Output struct {
	AudioHijackStyleFile string `yaml:"audioHijack"`
	JSONStyleFile        string `yaml:"json"`
	OBSStyleFileTemplate string `yaml:"obs"`
}

func getConfig() Config {
	filePath := "config/config.yaml"

	if !FileExists(filePath) {
		log.Fatal("ERROR: valid config/config.yaml is required")
	}

	yamlFile, err := ioutil.ReadFile(filePath)

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}
