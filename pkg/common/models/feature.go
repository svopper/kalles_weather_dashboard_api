package models

type Feature struct {
	Geometry   Geometry   `json:"geometry"`
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}
