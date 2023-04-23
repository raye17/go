package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Debug().
		Str("Scale", "833 cents").
		Float64("interval", 833.09).
		Msg("Fibonacci is here")
	log.Debug().Str("name", "tom").Send()
	log.Print("hello,world")
	log.Info().Msg("hello")
	log.Warn().Msg("warn")
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Layout})
	log.Info().Str("foo", "bar").Msg("hello,world")
	log.Warn().Str("test", "ts").Msg("warn:warn!")
}
