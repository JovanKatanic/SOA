package model

type TourKeyPoints struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Image          string  `json:"image"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	TourId         uint    `json:"tour_id"`
	Secret         string  `json:"secret"`
	PositionInTour int     `json:"position_in_tour"`
	PublicPointId  int     `json:"public_point_id"`
}
