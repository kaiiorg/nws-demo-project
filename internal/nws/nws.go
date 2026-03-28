package nws

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/kaiiorg/nws-demo-project/internal/nws/models"

	"github.com/rs/zerolog/log"
)

const (
	baseUrl = "https://api.weather.gov"
)

var (
	NwsError      = errors.New("nws error")
	NoNwsForecast = errors.New("nws returned a response, but it did not contain any forecasts")
)

func TempForCoords(lat, long float32) (*models.Forecast, int, error) {
	forecastUrl, nwsStatusCode, err := nwsGridSquare(lat, long)
	if err != nil {
		// Already logged details
		return nil, nwsStatusCode, err
	}

	log.Info().Str("forecastUrl", forecastUrl).Msg("got NWS grid square with forecast URL")

	f, nwsStatusCode, err := nwsTemp(forecastUrl)
	if err != nil {
		// Already logged details
		return nil, nwsStatusCode, err
	}

	return f, 0, nil
}

// nwsGridSquare requests to the NWS's /points/{latitude},{longitude} API endpoint for the given lat/long
// then returns the request URL needed for a forecast at that location or an error
func nwsGridSquare(lat, long float32) (string, int, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/points/%f,%f", baseUrl, lat, long), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to build request")
		return "", 0, err
	}

	// See authentication section at https://www.weather.gov/documentation/services-web-api
	req.Header.Set("User-Agent", "(nws-example, k349@live.com)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Float32("lat", lat).Float32("long", long).Msg("failed to request grid of provided lat/long")
		return "", 0, err
	}
	defer resp.Body.Close()

	// Read the body now. What we parse it into changes depending on the status code.
	// I understand that sometimes responses can be very large and thus may need to be streamed in chunks instead of loading them into memory all at once.
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Float32("lat", lat).Float32("long", long).Msg("failed to read body while requesting grid of provided lat/long")
		return "", 0, err
	}

	// NWS didn't like our request; attempt to parse the body as an error so we have some more detail on what went wrong, then return
	if resp.StatusCode != 200 {
		// Expect some sort of error message directly from NWS
		nwsErr := &models.Error{}
		_ = json.Unmarshal(rawBody, nwsErr)

		// We hopefully parsed a body, worst case, the NWS error title/detail will be empty
		log.Error().
			Float32("lat", lat).Float32("long", long).
			Int("code", resp.StatusCode).
			Str("nwsErrorTitle", nwsErr.Title).
			Str("nwsErrorDetail", nwsErr.Detail).
			Msg("non-200 status code")

		return "", resp.StatusCode, errors.Join(NwsError, errors.New(nwsErr.Title))
	}

	// We successfully got a grid square, we now need to parse out the predefined forecast request URL in the response
	p := &models.PointsResponse{}
	err = json.Unmarshal(rawBody, p)
	if err != nil {
		log.Error().
			Err(err).
			Float32("lat", lat).Float32("long", long).
			Msg("failed to parse body response while requesting grid of provided lat/long")
		// Logging the raw body would be a good thing to do here for a real application having an issue, but need to keep in mind that bodies may contain sensitive info and large bodies may stress our log aggregation
		return "", 0, err
	}

	return p.Properties.ForecastUrl, 200, nil
}

// nwsTemp calls the NWS API using the given forecastUrl and returns the temperature at that location or an error
// get forecastUrl from nwsGridSquare
func nwsTemp(forecastUrl string) (*models.Forecast, int, error) {
	req, err := http.NewRequest(http.MethodGet, forecastUrl, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to build request")
		return nil, 0, err
	}

	// See authentication section at https://www.weather.gov/documentation/services-web-api
	req.Header.Set("User-Agent", "(nws-example, k349@live.com)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Str("forecastUrl", forecastUrl).Msg("failed to request forecast at provided URL")
		return nil, 0, err
	}
	defer resp.Body.Close()

	// Read the body now. What we parse it into changes depending on the status code.
	// I understand that sometimes responses can be very large and thus may need to be streamed in chunks instead of loading them into memory all at once.
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Str("forecastUrl", forecastUrl).Msg("failed to read body from provided URL")
		return nil, 0, err
	}

	// NWS didn't like our request; attempt to parse the body as an error so we have some more detail on what went wrong, then return
	if resp.StatusCode != 200 {
		// Expect some sort of error message directly from NWS
		nwsErr := &models.Error{}
		_ = json.Unmarshal(rawBody, nwsErr)

		// We hopefully parsed a body, worst case, the NWS error title/detail will be empty
		log.Error().
			Str("forecastUrl", forecastUrl).
			Int("code", resp.StatusCode).
			Str("nwsErrorTitle", nwsErr.Title).
			Str("nwsErrorDetail", nwsErr.Detail).
			Msg("non-200 status code")

		return nil, resp.StatusCode, errors.Join(NwsError, errors.New(nwsErr.Title))
	}

	f := &models.ForecastResponse{}
	err = json.Unmarshal(rawBody, f)
	if err != nil {
		log.Error().
			Err(err).
			Str("forecastUrl", forecastUrl).
			Msg("failed to parse body response from provided URL")
		// Logging the raw body would be a good thing to do here for a real application having an issue, but need to keep in mind that bodies may contain sensitive info and large bodies may stress our log aggregation
		return nil, 0, err
	}

	if len(f.Properties.Periods) == 0 {
		return nil, 0, NoNwsForecast
	}

	// Assuming that the earliest forecast is always first; we could also loop through all of them and find the eariest one if needed
	// This also keeps a reference on the forecast and prevents the garbage collector from freeing f, but this is fine For Now™ because the API will be using it to build its own response and release it shortly
	return &f.Properties.Periods[0], 0, nil
}
