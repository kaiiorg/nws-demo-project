package characterize

import (
	"github.com/kaiiorg/nws-demo-project/internal/config"
)

type Characterize struct{}

func (ch *Characterize) Characterize(temp int, c *config.Forecast) string {
	if c == nil {
		panic("given nil config")
	}

	switch {
	case temp > c.Hot:
		return "hot"
	case temp < c.Cold:
		return "cold"
	default:
		return "moderate"
	}
}
