package config

type Forecast struct {
	// Hot anything above this value is considered hot, exclusive. Anything between this value and Cold is considered moderate.
	Hot int `json:"hot"`
	// Cold anything below this value is considered cold, exclusive. Anything between this value and Hot is considered moderate.
	Cold int `json:"cold"`
}
