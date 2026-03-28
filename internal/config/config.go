package config

import (
	"github.com/rs/zerolog/log"
)

var (
	defaultHot  = 90
	defaultCold = 60
)

type Config struct {
	Api      Api      `json:"api"`
	Forecast Forecast `json:"forecast"`
}

func LoadConfig(path string) *Config {
	c := &Config{}

	// We could load values from a config file from here, if desired. Not doing it now due to scope.
	if path != "" {
		log.Warn().Msg("loading config from a file is out of scope")
	}

	// Default to sane forecast mappings
	c.Forecast.Hot = defaultHot
	c.Forecast.Cold = defaultCold

	return c
}
