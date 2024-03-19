package repository

import (
	"tours_service/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) CreateTour(tour *model.Tour) (*model.Tour, error) {
	var maxID int
	if err := repo.DatabaseConnection.Table(`tours."Tour"`).Select("COALESCE(MAX(\"Id\"), 0)").Row().Scan(&maxID); err != nil {
		return nil, err
	}

	tour.ID = maxID + 1

	println(tour.ID)
	dbResult := repo.DatabaseConnection.Table(`tours."Tour"`).Create(tour)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	println("Rows affected: ", dbResult.RowsAffected)
	return tour, nil
}
