package models

type StationResponse struct {
	Stations []Station `json:"stations"`
}

type Station struct {
	Id   int    `json:"id"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
	Name string `json:"name"`
}
