package api

import (
	"github.com/kaiiorg/nws-demo-project/internal/config"
	nwsModels "github.com/kaiiorg/nws-demo-project/internal/nws/models"
)

type NwsClient interface {
	TempForCoords(lat, long float32) (*nwsModels.Forecast, int, error)
}

type Characterizer interface {
	Characterize(temp int, c *config.Forecast) string
}
