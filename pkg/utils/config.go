package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Playlists []string `yaml:"playlists"`
}

func processError(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func ReadConfiguration(cfg *Configuration) {
	f, err := os.Open(configurationFilePath)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
