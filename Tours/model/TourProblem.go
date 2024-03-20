package model

import "time"

type TourProblemCategory int
type TourProblemPriority int

type TourProblem struct {
	ID              int                 `json:"id" gorm:"column:Id; primaryKey"`
	TouristId       int                 `json:"touristId" gorm:"column:TouristId"`
	TourId          int                 `json:"tourId" gorm:"column:TourId"`
	Category        TourProblemCategory `json:"category" gorm:"column:Category"`
	Priority        TourProblemPriority `json:"priority" gorm:"column:Priority"`
	Description     string              `json:"description" gorm:"column:Description"`
	Time            time.Time           `json:"time" gorm:"column:Time"`
	IsSolved        bool                `json:"isSolved" gorm:"column:IsSolved"`
	Messages        TourProblemMessages `json:"messages" gorm:"type:jsonb; column:Messages"`
	Deadline        *time.Time          `json:"deadline" gorm:"column:Deadline"`
	AuthorUsername  string              `json:"authorUsername"`
	TouristUsername string              `json:"touristUsername"`
}

const (
	BOOKING TourProblemCategory = iota
	ITINERARY
	PAYMNET
	TRANSPORTATION
	GUIDE_SERVICES
	OTHER
)

const (
	LOW TourProblemPriority = iota
	MEDIUM
	HIGH
)
