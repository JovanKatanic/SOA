package repository

import (
	"tours_service/model"

	"gorm.io/gorm"
)

type TourProblemRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourProblemRepository) GetAll() (*[]model.TourProblem, error) {
	var allTourProblems []model.TourProblem
	dbResult := repo.DatabaseConnection.Table(`tours."TourProblems"`).Find(&allTourProblems)
	if dbResult != nil {
		return &allTourProblems, dbResult.Error
	}

	return &allTourProblems, nil
}

func (repo *TourProblemRepository) GetByTourId(tourId *int) (*model.TourProblem, error) {
	allTourProblems, err := repo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, value := range *allTourProblems {
		if value.TourId == *tourId {
			return &value, nil
		}
	}

	return nil, nil
}
