package stations

type DMIStationResponse struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	Name      string `json:"name"`
	StationID string `json:"stationId"`
}
