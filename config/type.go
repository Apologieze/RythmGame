package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	CurrentConfig  Config
	CurrentSetting Settings
)

type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	WindowFullscreen         bool
}

type Settings struct {
}

func Get(fileName string) {

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error when opening config file: ", err)
	}

	err = json.Unmarshal(content, &CurrentConfig)
	if err != nil {
		log.Fatal("Error when reading config file: ", err)
	}
}
