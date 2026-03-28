package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kaiiorg/nws-demo-project/internal/api/models"
	"github.com/kaiiorg/nws-demo-project/internal/config"
	"github.com/kaiiorg/nws-demo-project/internal/nws"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
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
	// Parse lat/long

	// Sanity check lat/long

	forecast, nwsStatusCode, err := nws.TempForCoords(36.7158451, -91.8739187)
	if err != nil {
		// Details already logged, we just need to respond to the HTTP request
		writeError(nwsStatusCode, err, w)
		return
	}

	log.Info().
		Int("temperature", forecast.Temperature).
		Str("shortForecast", forecast.ShortForecast).
		Msg("got forecast")

	// Characterize temp

	w.Write([]byte("forecast\n"))
}

func writeError(code int, err error, w http.ResponseWriter) {
	e := models.Error{
		Error: err.Error(),
	}

	rawBody, err := json.Marshal(e)
	if err != nil {
		// This should never happen
		panic(err)
	}

	if code == 0 {
		code = 500
	}
	w.WriteHeader(code)
	_, err = w.Write(rawBody)
	if err != nil {
		// This should never happen
		panic(err)
	}
}
