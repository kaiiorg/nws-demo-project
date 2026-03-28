package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kaiiorg/nws-demo-project/internal/api/models"
	"github.com/kaiiorg/nws-demo-project/internal/characterize"
	"github.com/kaiiorg/nws-demo-project/internal/config"
	"github.com/kaiiorg/nws-demo-project/internal/nws"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

var (
	forecastConfig    *config.Forecast
	nwsClient         NwsClient     = &nws.NwsClient{}
	tempCharacterizer Characterizer = &characterize.Characterize{}
)

func Run(c *config.Config) {
	// There's probably a better way of doing this with contexts, but this is simple enough for now
	forecastConfig = &c.Forecast

	log.Info().Uint16("port", c.Api.PortOrDefault()).Msg("API starting")

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

	forecast, nwsStatusCode, err := nwsClient.TempForCoords(36.7158451, -91.8739187)
	if err != nil {
		// Details already logged, we just need to respond to the HTTP request
		writeError(nwsStatusCode, err, w)
		return
	}

	// Characterize temp
	characterizedTemp := tempCharacterizer.Characterize(forecast.Temperature, forecastConfig)

	result := &models.ForecastResponse{
		Forecast:          forecast.ShortForecast,
		Short:             characterizedTemp,
		TemperatureFormat: "F",
		Temperature:       forecast.Temperature,
	}

	log.Info().
		Int("temperature", result.Temperature).
		Str("temperatureFormat", result.TemperatureFormat).
		Str("short", result.Short).
		Str("forecast", result.Forecast).
		Msg("got forecast")

	writeResponse(200, result, w)
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

func writeResponse(code int, body interface{}, w http.ResponseWriter) {
	rawBody, err := json.Marshal(body)
	if err != nil {
		// This should never happen
		panic(err)
	}

	if code == 0 {
		code = 200
	}
	w.WriteHeader(code)
	_, err = w.Write(rawBody)
	if err != nil {
		// This should never happen
		panic(err)
	}
}
