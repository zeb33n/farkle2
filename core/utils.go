package core

import (
	"encoding/json"
	"log"
	"os"
)

type TourType int

const (
	ROUNDROBIN TourType = iota
	ELIMINATION
)

type Config struct {
	FinalScore int
	TourType   TourType
	Bots       []string
	FirstTo    int
}

func (c *Config) LoadConfig(file string) {
	jsonFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Cant read config file", err)
	}
	err = json.Unmarshal(jsonFile, c)
	if err != nil {
		log.Fatal("Cant read config file", err)
	}
}

func WaitForKeyPress(verbose bool) string {
	b := make([]byte, 1)
	os.Stdin.Read(b)
	s := string(b)
	if verbose {
		if s != "\n" {
			print(s)
		}
	}
	return s
}

func nextSquare(i int) int {
	ui := uint32(i)
	// bit twiddling fun
	ui--
	ui |= ui >> 1
	ui |= ui >> 2
	ui |= ui >> 4
	ui |= ui >> 8
	ui |= ui >> 16
	ui++
	return int(ui)
}
