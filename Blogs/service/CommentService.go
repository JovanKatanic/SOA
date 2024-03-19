package service

import (
	"blogs_service/model"
	"blogs_service/repository"
	"time"
)

type CommentService struct {
	CommentRepository *repository.CommentRepository
}

func (service *CommentService) CreateComment(comment *model.Comment) error {
	comment.CreationDate = time.Now()
	err := service.CommentRepository.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}
