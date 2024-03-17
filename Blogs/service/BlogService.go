package service

import (
	"blogs_service/model"
	"blogs_service/repository"
	"time"
)

type BlogService struct {
	BlogRepository *repository.BlogRepository
}

func (service *BlogService) Create(blog *model.BlogPage) error {
	blog.CreationDate = time.Now()
	err := service.BlogRepository.CreateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}
