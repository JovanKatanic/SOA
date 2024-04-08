package model

import (
	"encoding/json"
	"io"
)

type Facility struct {
	ID            int     `json:"id" bson:"_id,omitempty"`
	Name          string  `json:"name" bson:"name"`
	Description   string  `json:"description" bson:"description,omitempty"`
	Image         string  `json:"image" bson:"image,omitempty"`
	Category      int     `json:"category" bson:"category,omitempty"`
	Latitude      float64 `json:"latitude" bson:"latitude,omitempty"`
	Longitude     float64 `json:"longitude" bson:"longitude,omitempty"`
	Discriminator string  `json:"discriminator" bson:"discriminator,omitempty"`
	Status        int     `json:"status" bson:"status,omitempty"`
	CreatorID     int     `json:"creator_id" bson:"creator_id,omitempty"`
}

func (p *Facility) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
