package config

type Forecast struct {
	// Exclusive, leave undefined define this as the top range value
	Max *int `json:"max,omitempty"`

	// Inclusive, leave undefined define this as the bottom range value
	Min *int `json:"min,omitempty"`
}
