package models

type PointsResponse struct {
	Properties PointsProperties `json:"properties"`
	// Other fields ignored because they aren't needed
}

type PointsProperties struct {
	ForecastUrl string `json:"forecast"`
	// Other fields ignored because they aren't needed
}
