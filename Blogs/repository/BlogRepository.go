package repository

import (
	"blogs_service/model"

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
