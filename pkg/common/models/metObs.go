package models

type MetObservationResponse struct {
	Date                    string                   `json:"date"`
	TemperatureObservations []TemperatureObservation `json:"temperatureObservations"`
	MaxAverage              float64                  `json:"maxAverage"`
	MinAverage              float64                  `json:"minAverage"`
}

type TemperatureObservation struct {
	Year int     `json:"year"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
}
