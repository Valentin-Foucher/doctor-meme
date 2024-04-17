package utils

import (
	"bytes"
	"log"
	"os"

	"github.com/Valentin-Foucher/doctor-meme/configs"
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
	decoder := yaml.NewDecoder(bytes.NewReader(configs.Config))
	err := decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
