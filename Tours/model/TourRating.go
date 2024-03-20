package model

import (
	"time"

	"github.com/lib/pq"
)

type TourRating struct {
	ID               int64          `json:"id" gorm:"column:Id; primaryKey"`
	PersonId         int16          `json:"personId" gorm:"column:PersonId"`
	TourId           int64          `json:"tourId" gorm:"column:TourId"`
	Mark             int32          `json:"mark" gorm:"column:Mark"`
	Comment          string         `json:"comment" gorm:"column:Comment"`
	DateOfVisit      time.Time      `json:"dateOfVisit" gorm:"column:DateOfVisit"`
	DateOfCommenting time.Time      `json:"dateOfCommenting" gorm:"column:DateOfCommenting"`
	Images           pq.StringArray `json:"images" gorm:"column:Images;type:string[]"`
}
