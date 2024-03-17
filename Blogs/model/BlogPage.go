package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type BlogPage struct {
	Id           int       `json:"id" gorm:"column:Id"`
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
		*r = Ratings{} // Postavljamo na prazan niz ako je prazan niz
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, &r)
}

/*func (r Ratings) Value() (driver.Value, error) {
	ratings := []Ratings{r}
	if len(ratings) == 0 { // Provjera da li je prazan niz
		return []byte("[]"), nil // Ako je prazan, vraćamo prazan JSON niz
	}
	return json.Marshal(ratings) // Inače, normalno serijalizujemo u JSON
}*/

func (r Ratings) Value() (driver.Value, error) {
	// Serijalizujemo Ratings u JSON format
	jsonData, err := json.Marshal(r)
	if jsonData == nil {
		return "[]", nil
	}
	if err != nil {
		return nil, err
	}

	// Vraćamo serijalizovani JSON kao string
	return "[" + string(jsonData) + "]", nil
}
