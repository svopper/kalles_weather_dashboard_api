package models

type DMIObservation struct {
	Type           string    `json:"type"`
	Features       []Feature `json:"features"`
	TimeStamp      string    `json:"timeStamp"`
	NumberReturned int64     `json:"numberReturned"`
	Links          []Link    `json:"links"`
}
