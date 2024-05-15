package model

type Wallet struct {
	ID      int    `json:"id" gorm:"column:Id;primaryKey"`
	UserId  int    `json:"userId" gorm:"column:UserId"`
	Balance string `json:"balance" gorm:"column:Balance"`
}
