package main

import (
	"flag"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	logLevel = flag.String("log-level", "info", "Set default Zerolog log level: trace, debug, info, warn, error, panic, etc")
)

func configureLogging() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	globalZerologLevel, err := zerolog.ParseLevel(*logLevel)
	if err != nil || globalZerologLevel == zerolog.NoLevel {
		log.Warn().
			Str("configuredLogLevel", *logLevel).
			Msgf("Log level set to unexpected level; defaulting to %s level", zerolog.InfoLevel.String())
		globalZerologLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(globalZerologLevel)
}
