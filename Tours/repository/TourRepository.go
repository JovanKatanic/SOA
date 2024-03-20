package repository

import (
	"tours_service/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) CreateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Table(`tours."Tour"`).Create(tour)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) UpdateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Table(`tours."Tour"`).Save(tour)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) FindByID(id int) (*model.Tour, error) {
	var tour model.Tour
	if err := repo.DatabaseConnection.Table(`tours."Tour"`).First(&tour, id).Error; err != nil {
		return nil, err
	}
	var keypoints []model.Keypoint
	if err := repo.DatabaseConnection.Table(`tours."TourKeyPoints"`).Where(`"tours"."TourKeyPoints"."TourId" = ?`, tour.ID).Find(&keypoints).Error; err != nil {
		return nil, err
	}
	tour.KeyPoints = keypoints

	//fmt.Println(tour)
	return &tour, nil
}
