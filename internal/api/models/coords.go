package models

type Coords struct {
	// Latitude is the WGS84 latitude (X)
	Latitude float32 `json:"latitude"`
	// Longitude is the WGS84 longitude (Y)
	Longitude float32 `json:"longitude"`
}
