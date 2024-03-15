package model

type Keypoint struct {
	Id             int     `json:"id" gorm:"column:Id"`
	Name           string  `json:"name" gorm:"column:Name"`
	Description    string  `json:"description" gorm:"column:Description"`
	Image          string  `json:"image" gorm:"column:Image"`
	Latitude       float64 `json:"latitude" gorm:"column:Latitude"`
	Longitude      float64 `json:"longitude" gorm:"column:Longitude"`
	TourId         int     `json:"tourId,omitempty" gorm:"column:TourId"`
	PositionInTour int     `json:"positionInTour,omitempty" gorm:"column:PositionInTour"`
	PublicPointId  int     `json:"publicPointId,omitempty" gorm:"column:PublicPointId"`
	Secret         string  `json:"secret,omitempty" gorm:"column:Secret"`
	Discriminator  string  `json:"discriminator" gorm:"column:Discriminator"`
}
