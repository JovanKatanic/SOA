package model

import (
	"encoding/json"
	"io"
)

type Keypoint struct {
	ID             int     `json:"id" bson:"_id,omitempty"`
	Name           string  `json:"name" bson:"name,omitempty"`
	Description    string  `json:"description" bson:"description,omitempty"`
	Image          string  `json:"image" bson:"image,omitempty"`
	Latitude       float64 `json:"latitude" bson:"latitude,omitempty"`
	Longitude      float64 `json:"longitude" bson:"longitude,omitempty"`
	TourId         int     `json:"tourId,omitempty" bson:"tourId,omitempty"`
	PositionInTour int     `json:"positionInTour,omitempty" bson:"positionInTour,omitempty"`
	Secret         string  `json:"secret,omitempty" bson:"secret,omitempty"`
	Discriminator  string  `json:"discriminator" bson:"discriminator,omitempty"`
	//PublicPointId  int     `json:"publicPointId,omitempty" gorm:"column:PublicPointId"`
}

func (p *Keypoint) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
