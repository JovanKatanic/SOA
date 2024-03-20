package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TransportationType int

type TourDuration struct {
	TimeInSeconds  uint               `json:"TimeInSeconds"`
	Transportation TransportationType `json:"Transportation"`
}

const (
	Walking TransportationType = iota
	Bicycle
	Car
)

func (d TourDuration) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *TourDuration) Scan(value interface{}) error {
	if value == nil {
		*d = TourDuration{}
		return nil
	}

	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}

	return json.Unmarshal(bytes, d)
}

type TourDurations []TourDuration

func (td TourDurations) Value() (driver.Value, error) {
	return json.Marshal(td)
}

func (td *TourDurations) Scan(value interface{}) error {
	if value == nil {
		*td = TourDurations{}
		return nil
	}

	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}

	return json.Unmarshal(bytes, td)
}
