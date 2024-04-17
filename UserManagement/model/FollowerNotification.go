package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type FollowerNotification struct {
	Content       string     `json:"content" bson:"content"`
	TimeOfArrival *time.Time `json:"timeOfArrival" bson:"timeOfArrival"`
	Read          bool       `json:"read" bson:"read"`
}

func (n FollowerNotification) Value() (driver.Value, error) {
	return json.Marshal(n)
}

func (n *FollowerNotification) Scan(value interface{}) error {
	if value == nil {
		*n = FollowerNotification{}
		return nil
	}

	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}

	return json.Unmarshal(bytes, n)
}

type Notifications []FollowerNotification

func (n Notifications) Value() (driver.Value, error) {
	return json.Marshal(n)
}

func (n *Notifications) Scan(value interface{}) error {
	if value == nil {
		*n = Notifications{}
		return nil
	}

	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}

	return json.Unmarshal(bytes, n)
}
