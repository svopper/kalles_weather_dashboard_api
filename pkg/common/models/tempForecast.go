package models

type TemperatureForecast struct {
	Date        string  `json:"date"`
	Temperature float64 `json:"temperature"`
}
