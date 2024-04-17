package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type HttpClientConfig struct {
	BaseUrl string `yaml:"base_url"`
}

type HttpConfig struct {
	MaxRetries int              `yaml:"max_retries"`
	Youtube    HttpClientConfig `yaml:"youtube"`
}

type YoutubeConfig struct {
	PlaylistIDs []string `yaml:"playlists"`
}

type Configuration struct {
	YoutubeConfig YoutubeConfig `yaml:"youtube"`
	HttpConfig    HttpConfig    `yaml:"http_client"`
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
