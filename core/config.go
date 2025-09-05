package core

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	BotNames   []string
	FinalScore int
	FirstTo    int
}

const CONFIGFILE string = "config.json"

var CONFIG Config = Config{
	BotNames:   []string{},
	FinalScore: 10000,
	FirstTo:    3,
}

func (c *Config) LoadConfig(file string) {
	jsonFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Cant read config file", err)
	}
	err = json.Unmarshal(jsonFile, c)
	if err != nil {
		log.Fatal("Error parsing config file", err)
	}
}
