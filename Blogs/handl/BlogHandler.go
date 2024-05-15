package handl

import (
	"blogs_service/model"
	"blogs_service/proto/blogs"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type BlogHandler struct {
	blogs.UnimplementedBlogServiceServer
	DatabaseConnection *gorm.DB
}

func (h BlogHandler) GetBlog(ctx context.Context, request *blogs.GetBlogRequest) (*blogs.GetBlogResponse, error) {
	var blog model.BlogPage
	fmt.Println("Usao u handler")
	if err := h.DatabaseConnection.Table(`blog."Blogs"`).Where(`"Blogs"."Id" = ?`, request.Id).First(&blog).Error; err != nil {
		return nil, err
	}

	blogic := &blogs.Blog{
		Id:          1,
		Title:       blog.Title,
		Description: blog.Description,
		Status:      1,
		UserId:      1,
	}

	fmt.Println(blog)

	response := &blogs.GetBlogResponse{
		Blog: blogic,
	}

	return response, nil
}
