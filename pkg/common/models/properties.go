package models

type Properties struct {
	Created     string  `json:"created"`
	Observed    string  `json:"observed"`
	ParameterID string  `json:"parameterId"`
	StationID   string  `json:"stationId"`
	Value       float64 `json:"value"`
	Name        string  `json:"name"`
	Temperature float64 `json:"temperature-2m"`
	Step        string  `json:"step"`
}
