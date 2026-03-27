package api

import (
	"net/http"
	"fmt"

	"github.com/kaiiorg/nws-demo-project/internal/config"

	"github.com/rs/zerolog/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(c *config.Config) {
	log.Info().
		Uint16("port", c.Api.PortOrDefault()).
		Msg("hello world from api")
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/v1/forecast", getForecast)

	err := http.ListenAndServe(
		fmt.Sprintf(":%d", c.Api.PortOrDefault()),
		r,
	)
	if err != nil {
		log.Error().Err(err).Msg("Stopped listening for HTTP traffic")
	}
}


func getForecast(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("forecast\n"))
}
