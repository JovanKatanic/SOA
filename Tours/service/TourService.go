package service

import (
	"tours_service/model"
	"tours_service/repository"
)

type TourService struct {
	TourRepository *repository.TourRepository
}

func (service *TourService) CreateTour(tour *model.Tour) error {
	err := service.TourRepository.CreateTour(tour)

	if err != nil {
		return err
	}
	return nil
}
