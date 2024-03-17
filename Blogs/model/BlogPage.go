package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type BlogPage struct {
	Id           int       `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	Title        string    `json:"title" gorm:"column:Title"`
	Description  string    `json:"description" gorm:"column:Description"`
	CreationDate time.Time `json:"creationDate" gorm:"column:CreationDate"`
	Status       uint      `json:"status" gorm:"column:Status"`
	UserId       int       `json:"userId" gorm:"column:UserId"`
	RatingSum    int       `json:"ratingSum" gorm:"column:RatingSum"`
	Ratings      []Ratings `json:"ratings" gorm:"type:jsonb;column:Ratings;"`
}

type Ratings struct {
	UserId       int       `json:"userId" gorm:"column:UserId"`
	CreationDate time.Time `json:"creationDate" gorm:"column:CreationDate"`
	RatingValue  int       `json:"ratingValue" gorm:"column:RatingValue"`
}

func (r *Ratings) Scan(value interface{}) error {
	if value == nil {
		*r = Ratings{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, r)
}
func (r Ratings) Value() (driver.Value, error) {
	if r.RatingValue == 0 {
		ratings := make([]Ratings, 0)
		return json.Marshal(ratings)
	}
	ratings := []Ratings{r}
	return json.Marshal(ratings)
	//return json.Marshal(r)
}
