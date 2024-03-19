package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Tour struct {
	ID            int            `json:"id" gorm:"column:Id"`
	Name          string         `json:"name" gorm:"column:Name"`
	Description   string         `json:"description" gorm:"column:Description"`
	Difficulty    int            `json:"difficulty" gorm:"column:Difficulty"`
	Tags          pq.StringArray `json:"tags" gorm:"type:string[];column:Tags"`
	Status        int            `json:"status" gorm:"column:Status"`
	Price         float64        `json:"price" gorm:"column:Price"`
	AuthorID      int            `json:"authorId" gorm:"column:AuthorId"`
	Equipment     pq.Int64Array  `json:"equipment" gorm:"type:integer[];column:Equipment"`
	DistanceInKm  float64        `json:"distanceInKm" gorm:"column:DistanceInKm"`
	ArchivedDate  *time.Time     `json:"archivedDate" gorm:"column:ArchivedDate"`
	PublishedDate *time.Time     `json:"publishedDate" gorm:"column:PublishedDate"`
	Durations     TourDurations  `json:"durations" gorm:"type:jsonb;column:Durations"`
	KeyPoints     []Keypoint     `json:"keyPoints" gorm:"-"`
	Image         string         `json:"image" gorm:"-"`
}
type TourDurations []TourDuration

func (r TourDurations) Value() (driver.Value, error) {
	return json.Marshal(r)
}

type TourDuration struct {
	TimeInSeconds  uint32 `json:"timeInSeconds" gorm:"column:TimeInSeconds"`
	Transportation int    `json:"transportation" gorm:"column:Transportation"`
}

// func (r *TourDuration) Scan(value interface{}) error {
// 	if value == nil {
// 		*r = TourDuration{}
// 		return nil
// 	}
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return fmt.Errorf("Scan source is not []byte")
// 	}
// 	return json.Unmarshal(bytes, r)
// }
