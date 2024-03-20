package repository

import (
	"tours_service/model"

	"gorm.io/gorm"
)

type TourRatingRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRatingRepository) CreateTourRating(tourRating *model.TourRating) error {
	dbResult := repo.DatabaseConnection.Table(`tours."TourRatings"`).Create(tourRating)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
