package config

const (
	defaultPort = 8080
)

type Api struct {
	// What port to host the API server on. Defaults to 8080
	Port uint16 `json:"Port"`
}

// Gets the configured Port or the default if the value is invalid
func (api *Api) PortOrDefault() uint16 {
	if api.Port == 0 {
		return defaultPort
	} else {
		return api.Port
	}
}
