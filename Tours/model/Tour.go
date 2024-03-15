package model

import "time"

type Tour struct {
	ID            int            `json:"id" gorm:"column:Id"`
	Name          string         `json:"name" gorm:"column:Name"`
	Description   string         `json:"description" gorm:"column:Description"`
	Difficulty    int            `json:"difficulty" gorm:"column:Difficulty"`
	Tags          []string       `json:"tags" gorm:"column:Tags"`
	Status        int            `json:"status" gorm:"column:Status"`
	Price         float64        `json:"price" gorm:"column:Price"`
	AuthorID      int            `json:"authorId" gorm:"column:AuthorId"`
	Equipment     []int          `json:"equipment" gorm:"column:Equipment"`
	DistanceInKm  float64        `json:"distanceInKm" gorm:"column:DistanceInKm"`
	ArchivedDate  *time.Time     `json:"archivedDate" gorm:"column:ArchivedDate"`
	PublishedDate *time.Time     `json:"publishedDate" gorm:"column:PublishedDate"`
	Durations     []TourDuration `json:"durations" gorm:"-"`
	KeyPoints     []Keypoint     `json:"keyPoints" gorm:"-"`
}

type TourDuration struct {
	TimeInSeconds  uint32 `json:"timeInSeconds" gorm:"column:TimeInSeconds"`
	Transportation int    `json:"transportation" gorm:"column:Transportation"`
}
