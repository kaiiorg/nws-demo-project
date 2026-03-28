package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/kaiiorg/nws-demo-project/internal/api/models"
	"github.com/kaiiorg/nws-demo-project/internal/characterizer"
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

var (
	ErrInvalidLatitude  = errors.New("invalid latitude")
	ErrInvalidLongitude = errors.New("invalid longitude")
)

func Run(c *config.Config) {
	// There's probably a better way of doing this with contexts, but this is simple enough for now
	forecastConfig = &c.Forecast

	log.Info().Uint16("port", c.Api.PortOrDefault()).Msg("API starting")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/v1/forecast", postForecast)

	err := http.ListenAndServe(
		fmt.Sprintf(":%d", c.Api.PortOrDefault()),
		r,
	)
	if err != nil {
		log.Error().Err(err).Msg("Stopped listening for HTTP traffic")
	}
}

func postForecast(w http.ResponseWriter, r *http.Request) {
	// Read request body
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read request body")
		writeError(500, err, w)
		return
	}

	// Parse lat/long
	coords := &models.Coords{}
	err = json.Unmarshal(rawBody, coords)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse request body")
		// Logging the raw body would be a good thing to do here for a real application having an issue, but need to keep in mind that bodies may contain sensitive info and large bodies may stress our log aggregation, not to mention security
		writeError(500, err, w)
		return
	}

	// Sanity check lat/long
	err = latLongSanityCheck(coords)
	if err != nil {
		log.Error().Err(err).Msg("given invalid coordinates")
		// Logging the raw body would be a good thing to do here for a real application having an issue, but need to keep in mind that bodies may contain sensitive info and large bodies may stress our log aggregation, not to mention security
		writeError(400, err, w)
		return
	}

	forecast, nwsStatusCode, err := nwsClient.TempForCoords(coords.Latitude, coords.Longitude)
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

func latLongSanityCheck(coords *models.Coords) error {
	switch {
	case coords.Latitude < -90.0, coords.Latitude > 90.0:
		return ErrInvalidLatitude
	case coords.Longitude < -180.0, coords.Longitude > 180:
		return ErrInvalidLongitude
	}
	return nil
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
