package repository

import (
	"tours_service/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) CreateTour(tour *model.Tour) (*model.Tour, error) {
	dbResult := repo.DatabaseConnection.Table(`tours."Tour"`).Create(tour)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	println("Rows affected: ", dbResult.RowsAffected)
	return tour, nil
}

func (repo *TourRepository) UpdateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Table(`tours."Tour"`).Save(tour)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
