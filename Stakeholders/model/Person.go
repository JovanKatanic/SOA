package model

type Person struct {
	Id         int     `gorm:"column:Id"`
	UserId     int     `gorm:"column:UserId"`
	Name       string  `gorm:"column:Name"`
	Surname    string  `gorm:"column:Surname"`
	Email      string  `gorm:"column:Email"`
	ProfilePic string  `gorm:"column:ProfilePic;default:null"`
	Biography  string  `gorm:"column:Biography;default:null"`
	Motto      string  `gorm:"column:Motto;default:null"`
	Latitude   float64 `gorm:"column:Latitude;default:null"`
	Longitude  float64 `gorm:"column:Longitude;default:null"`
}
