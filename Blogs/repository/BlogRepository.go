package repository

import (
	"blogs_service/model"
	"fmt"

	"gorm.io/gorm"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogRepository) CreateBlog(blog *model.BlogPage) error {

	dbResult := repo.DatabaseConnection.Table(`blog."Blogs"`).Create(blog)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *BlogRepository) GetAll() ([]model.BlogPage, error) {
	var blogs []model.BlogPage
	if err := repo.DatabaseConnection.Table(`blog."Blogs"`).Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (repo *BlogRepository) FindByID(id int) (*model.BlogPage, error) {
	var blog model.BlogPage
	if err := repo.DatabaseConnection.Table(`blog."Blogs"`).First(&blog, id).Error; err != nil {
		return nil, err
	}
	fmt.Println(blog)
	return &blog, nil
}

func (repo *BlogRepository) UpdateOneBlog(blog *model.BlogPage) error {
	dbResult := repo.DatabaseConnection.Table(`blog."Blogs"`).Save(blog)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
