package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type ss struct {
}

func (s ss) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	fmt.Println("llllllllllll")
}
func main() {
	file, _ := os.Create("log.md")
	logger := log.Output(file)
	logger.Info().Str("foo", "bar").Msg("hello world")
}
