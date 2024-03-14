package model

type User struct {
	Id                     int    `gorm:"column:Id"`
	Username               string `gorm:"column:Username"`
	Password               string `gorm:"column:Password"`
	Role                   uint   `gorm:"column:Role"`
	IsActive               bool   `gorm:"column:IsActive"`
	ResetPasswordToken     string `gorm:"column:ResetPasswordToken;default:null"`
	EmailVerificationToken string `gorm:"column:EmailVerificationToken;default:null"`
}
