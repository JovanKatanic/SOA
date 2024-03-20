package repository

import (
	"tours_service/model"

	"gorm.io/gorm"
)

type KeypointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *KeypointRepository) CreateKeypoint(keypoint *model.Keypoint) error {
	dbResult := repo.DatabaseConnection.Table(`tours."TourKeyPoints"`).Create(keypoint)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
