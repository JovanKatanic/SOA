package repository

import (
	"fmt"
	"stakeholders_service/model"

	"gorm.io/gorm"
)

type UserRepostiroy struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserRepostiroy) CreateUser(user *model.User) error {
	dbResult := repo.DatabaseConnection.Table(`stakeholders."Users"`).Create(user)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *UserRepostiroy) FindByUsername(username string) error {
	user := model.User{}
	dbResult := repo.DatabaseConnection.Table(`stakeholders."Users"`).Find(&user, "\"Username\" = ?", username)
	fmt.Println(user)
	if dbResult.Error != nil {
		println("Greksa prilikom poziva upita prema bazi")
		return dbResult.Error
	}

	if user.Id != 0 {
		println("Postoji vec korisnik sa istim username")
		return fmt.Errorf("postoji vec korisnik sa istim username")
	}

	return nil
}
