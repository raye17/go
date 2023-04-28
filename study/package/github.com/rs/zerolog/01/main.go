package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger.Info().Str("foo", "bar").Msg("hello world")
}
