package models

type ForecastResponse struct {
	Properties ForecastProperies `json:"properties"`
}

type ForecastProperies struct {
	Periods []Forecast `json:"periods"`
}

type Forecast struct {
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"shortForecast"`
}
