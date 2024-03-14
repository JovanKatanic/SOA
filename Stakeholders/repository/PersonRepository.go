package repository

import (
	"stakeholders_service/model"

	"gorm.io/gorm"
)

type PersonRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *PersonRepository) CreatePerson(person *model.Person) error {
	dbResult := repo.DatabaseConnection.Table(`stakeholders."People"`).Create(person)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
