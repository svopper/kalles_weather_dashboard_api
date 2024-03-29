package models

type OceanObservationsResponse struct {
	Date         string             `json:"date"`
	Observations []OceanObservation `json:"observations"`
}

type OceanObservationResponse struct {
	Date        string           `json:"date"`
	Observation OceanObservation `json:"observation"`
}

type OceanObservation struct {
	StationId         int     `json:"stationId"`
	StationName       string  `json:"stationName"`
	MaxTemp24H        float64 `json:"maxTemp24H"`
	LatestTemperature float64 `json:"latestTemperature"`
}
