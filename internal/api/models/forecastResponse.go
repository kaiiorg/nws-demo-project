package models

type ForecastResponse struct {
	Forecast          string `json:"forecast"`
	Short             string `json:"short"`
	TemperatureFormat string `json:"temperatureFormat"`
	Temperature       int    `json:"temperature"`
}
