package config

type Forecast struct {
	Hot      ForecastCharacterization `json:"hot"`
	Moderate ForecastCharacterization `json:"moderate"`
	Cold     ForecastCharacterization `json:"cold"`
}

type ForecastCharacterization struct {
	Max int `json:"max,omitempty"`
	Min int `json:"min,omitempty"`
}
