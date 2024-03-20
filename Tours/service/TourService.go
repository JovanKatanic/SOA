package service

import (
	"tours_service/model"
	"tours_service/repository"
)

type TourService struct {
	TourRepository *repository.TourRepository
}

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepository.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) Update(tour *model.Tour) error {
	err := service.TourRepository.UpdateTour(tour)
	if err != nil {
		return err
	}
	return nil
}
func (s *TourService) FindByID(id int) (*model.Tour, error) {
	return s.TourRepository.FindByID(id)
}
