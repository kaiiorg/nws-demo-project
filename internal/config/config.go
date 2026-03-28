package config

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

	// TODO parse from JSON file

	// Default to sane forecast mappings if we failed to parse the JSON file
	c.Forecast.Hot = defaultHot
	c.Forecast.Cold = defaultCold

	return c
}
