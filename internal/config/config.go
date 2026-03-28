package config

var (
	defaultHotMin      = 90
	defaultColdMax     = 60
	defaultModerateMax = defaultHotMin
	defaultModerateMin = defaultColdMax
)

type Config struct {
	Api      Api      `json:"api"`
	Forecast Forecast `json:"forecast"`
}

func LoadConfig(path string) *Config {
	c := &Config{}

	// TODO parse from JSON file

	// Default to sane forecast mappings if we failed to parse the JSON file
	c.Forecast.Hot.Min = defaultHotMin
	c.Forecast.Cold.Max = defaultColdMax
	c.Forecast.Moderate.Max = defaultModerateMax
	c.Forecast.Moderate.Min = defaultModerateMin

	return c
}
