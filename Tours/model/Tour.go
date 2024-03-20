package model

import (
	"time"

	"github.com/lib/pq"
)

type TourDifficulty int
type TourStatus int

type Tour struct {
	ID            int             `json:"id" gorm:"column:Id; primaryKey;"`
	Name          string          `json:"name" gorm:"column:Name"`
	Description   string          `json:"description" gorm:"column:Description"`
	Difficulty    TourDifficulty  `json:"difficulty" gorm:"column:Difficulty"`
	Tags          pq.StringArray  `json:"tags" gorm:"type:text[]; column:Tags"`
	Status        int             `json:"status" gorm:"column:Status"`
	Price         float64         `json:"price" gorm:"column:Price"`
	AuthorId      int             `json:"authorId" gorm:"column:AuthorId"`
	Equipment     pq.Int32Array   `json:"equipment" gorm:"type:integer[]; column:Equipment"`
	DistanceInKm  float64         `json:"distanceInKm" gorm:"column:DistanceInKm"`
	ArchivedDate  time.Time       `json:"archivedDate" gorm:"column:ArchivedDate"`
	PublishedDate time.Time       `json:"publishedDate" gorm:"column:PublishedDate"`
	Durations     TourDurations   `json:"durations" gorm:"type:jsonb; column:Durations"`
	KeyPoints     []TourKeyPoints `json:"keyPoints"`
	Image         string          `json:"image" gorm:"column:Image"`
}

const (
	Beginner TourDifficulty = iota
	Intermediate
	Advanced
	Pro
)

const (
	Draft TourStatus = iota
	Published
	Archived
	TouristMade
)
