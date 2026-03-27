package config

var (
	defaultHotLabel = "hot"
	defaultHotMin   = 90

	defaultColdLabel = "cold"
	defaultColdMax   = 60

	defaultModerateLabel = "moderate"
	defaultModerateMax   = defaultHotMin
	defaultModerateMin   = defaultColdMax
)

type Config struct {
	Api      Api                 `json:"api"`
	Forecast map[string]Forecast `json:"forecast"`
}

func LoadConfig(path string) *Config {
	c := &Config{
		Forecast: make(map[string]Forecast),
	}

	// TODO parse from JSON file

	// Default to sane forecast mappings
	if len(c.Forecast) == 0 {
		c.Forecast[defaultHotLabel] = Forecast{
			Min: &defaultHotMin,
		}
		c.Forecast[defaultModerateLabel] = Forecast{
			Max: &defaultModerateMax,
			Min: &defaultModerateMin,
		}
		c.Forecast[defaultColdLabel] = Forecast{
			Max: &defaultColdMax,
		}
	}

	return c
}
