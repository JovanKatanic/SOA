package service

import (
	"tours_service/model"
	"tours_service/repository"
)

type FacilityService struct {
	FacilityRepository *repository.FacilityRepository
}

func (service *FacilityService) Create(facility *model.Facility) error {
	err := service.FacilityRepository.CreateFacility(facility)
	if err != nil {
		return err
	}
	return nil
}
