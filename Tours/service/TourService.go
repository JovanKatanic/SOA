package service

import (
	"fmt"
	"tours_service/model"
	"tours_service/repository"
)

type TourService struct {
	TourRepository     *repository.TourRepository
	KeypointRepository *repository.KeypointRepository
}

func (service *TourService) GetAll() (*[]model.Tour, error) {
	tours, err := service.TourRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return tours, nil
}

func (service *TourService) CreateTour(tour *model.Tour) error {

	err := service.TourRepository.Insert(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) UpdateTour(tour *model.Tour) error {
	err := service.TourRepository.Update(tour)
	if err != nil {
		return err
	}
	return nil
}
func (s *TourService) GetTourById(id int) (*model.Tour, error) {
	tour, _ := s.TourRepository.GetById(id)
	if tour == nil {
		// Handle case where tour is nil (not found)
		fmt.Println("Tour not found")
		return nil, nil
	}
	keypoints, err := s.KeypointRepository.GetByTourId(id)
	tour.KeyPoints = keypoints
	return tour, err
}
