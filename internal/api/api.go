package api

import (
	"github.com/kaiiorg/nws-demo-project/internal/config"

	"github.com/rs/zerolog/log"
)

func Run(c *config.Config) {
	log.Info().
		Uint16("port", c.Api.PortOrDefault()).
		Msg("hello world from api")
}
