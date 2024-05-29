package handl

import (
	"blogs_service/model"
	"blogs_service/proto/blogs"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BlogHandler struct {
	blogs.UnimplementedBlogServiceServer
	DatabaseConnection *gorm.DB
}

func (h BlogHandler) GetBlog(ctx context.Context, request *blogs.GetBlogRequest) (*blogs.GetBlogResponse, error) {
	var blog model.BlogPage
	if err := h.DatabaseConnection.Table(`blog."Blogs"`).Where(`"Blogs"."Id" = ?`, request.Id).First(&blog).Error; err != nil {
		return nil, err
	}

	ratingList := make([]*blogs.Rating, len(blog.Ratings))
	for i, rating := range blog.Ratings {
		ratingList[i] = &blogs.Rating{
			UserId:       int32(rating.UserId),
			CreationDate: rating.CreationDate.String(),
			RatingValue:  int32(rating.RatingValue),
		}
	}

	blogic := &blogs.Blog{
		Id:           int32(blog.Id),
		Title:        blog.Title,
		Description:  blog.Description,
		CreationDate: blog.CreationDate.String(),
		Status:       int32(blog.Status),
		UserId:       int32(blog.UserId),
		RatingSum:    int32(blog.RatingSum),
		Ratings:      ratingList,
	}

	fmt.Println(blog)

	response := &blogs.GetBlogResponse{
		Blog: blogic,
	}

	return response, nil
}

func (h BlogHandler) CreateBlog(ctx context.Context, request *blogs.Blog) (*blogs.GetBlogResponse, error) {

	var ratingsList model.BlogRatings
	for _, rating := range request.Ratings {
		layout := "2006-01-02T15:04:05.000Z"

		dateTime, _ := time.Parse(layout, rating.CreationDate)
		newRating := model.Rating{
			UserId:       int(rating.UserId),
			CreationDate: dateTime,
			RatingValue:  int(rating.RatingValue),
		}
		ratingsList = append(ratingsList, newRating)
	}
	var blog = model.BlogPage{
		Id:           int(request.Id),
		Title:        request.Title,
		Description:  request.Description,
		CreationDate: time.Now(),
		Status:       uint(request.Status),
		UserId:       int(request.UserId),
		RatingSum:    int(request.RatingSum),
		Ratings:      ratingsList,
	}

	if err := h.DatabaseConnection.Table(`blog."Blogs"`).Create(blog).Error; err != nil {
		return nil, err
	}

	response := &blogs.GetBlogResponse{
		Blog: request,
	}

	return response, nil
}
