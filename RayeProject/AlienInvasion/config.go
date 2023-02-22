package main

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	ScreenWidth  int        `json:"screenWidth"`
	ScreenHeight int        `json:"screenHeight"`
	Title        string     `json:"title"`
	BgColor      color.RGBA `json:"bgColor"`
}

func loadConfig() *Config {
	f, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("os.open failed: %v\n", err)
	}
	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.decode failed: %v\n", err)
	}
	return &cfg
}
