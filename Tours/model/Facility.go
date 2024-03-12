package model

type Facility struct {
	ID            uint    `json:"id" gorm:"column:Id"`
	Name          string  `json:"name" gorm:"column:Name"`
	Description   string  `json:"description" gorm:"column:Description"`
	Image         string  `json:"image" gorm:"column:Image"`
	Category      uint    `json:"category" gorm:"column:Category"`
	Latitude      float64 `json:"latitude" gorm:"column:Latitude"`
	Longitude     float64 `json:"longitude" gorm:"column:Longitude"`
	Discriminator string  `json:"discriminator" gorm:"column:Discriminator"`
	Status        uint    `json:"status" gorm:"column:Status"`
	CreatorID     uint    `json:"creator_id" gorm:"column:CreatorId"`
}
