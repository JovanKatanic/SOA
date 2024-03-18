package model

type TransportationType int

type TourDuration struct {
	ID             uint               `json:"id"`
	TimeInSeconds  uint               `json:"time_in_seconds"`
	Transportation TransportationType `json:"transportation"`
}

const (
	Walking TransportationType = iota
	Bicycle
	Car
)
