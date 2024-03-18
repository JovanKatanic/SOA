package model

import (
	"time"
)

type TourDifficulty int
type TourStatus int

type Tour struct {
	ID            uint            `json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Difficulty    TourDifficulty  `json:"difficulty"`
	Tags          []string        `json:"tags"`
	Price         float64         `json:"price"`
	AuthorId      uint            `json:"author_id"`
	Equipment     []uint          `json:"equipment"`
	DistanceInKm  float64         `json:"distance_in_km"`
	ArchivedDate  time.Time       `json:"archived_date"`
	PublishedDate time.Time       `json:"published_date"`
	Durations     []TourDuration  `json:"durations"`
	KeyPoints     []TourKeyPoints `json:"keypoints"`
	Image         string          `json:"image"`
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
