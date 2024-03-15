package service

import (
	"tours_service/model"
	"tours_service/repository"
)

type KeypointService struct {
	KeypointRepository *repository.KeypointRepository
}

func (service *KeypointService) Create(keypoint *model.Keypoint) error {
	err := service.KeypointRepository.CreateKeypoint(keypoint)
	if err != nil {
		return err
	}
	return nil
}
