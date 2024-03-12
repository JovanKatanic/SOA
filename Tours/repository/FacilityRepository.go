package repository

import (
	"tours_service/model"

	"gorm.io/gorm"
)

type FacilityRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *FacilityRepository) CreateFacility(facility *model.Facility) error {
	dbResult := repo.DatabaseConnection.Table(`tours."Facilities"`).Create(facility)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
