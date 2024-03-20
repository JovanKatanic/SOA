package service

import (
	"tours_service/model"
	"tours_service/repository"
)

type TourService struct {
	TourRepository *repository.TourRepository
}

func (service *TourService) GetAll() (*[]model.Tour, error) {
	tours, err := service.TourRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func (service *TourService) CreateTour(tour *model.Tour) (*model.Tour, error) {
	var createdTour *model.Tour

	createdTour, err := service.TourRepository.CreateTour(tour)

	if err != nil {
		return nil, err
	}

	return createdTour, nil
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
