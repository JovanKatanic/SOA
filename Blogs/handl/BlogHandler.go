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

func (h BlogHandler) GetBlog(ctx context.Context, request *blogs.GetBlogRequest) (*blogs.Blog, error) {
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

	return blogic, nil
}

func (h BlogHandler) CreateBlog(ctx context.Context, request *blogs.Blog) (*blogs.Blog, error) {
	fmt.Printf("Received Blog: %+v\n", request)
	fmt.Printf("Received Blog ID: %d, Title: %s, Description: %s, UserId: %d, RatingSum: %d, Ratings: %+v\n",
		request.Id, request.Title, request.Description, request.UserId, request.RatingSum, request.Ratings)

	var ratingsList model.BlogRatings
	for _, rating := range request.Ratings {
		layout := "2006-01-02T15:04:05.000Z"
		dateTime, err := time.Parse(layout, rating.CreationDate)
		if err != nil {
			return nil, err
		}
		newRating := model.Rating{
			UserId:       int(rating.UserId),
			CreationDate: dateTime,
			RatingValue:  int(rating.RatingValue),
		}
		ratingsList = append(ratingsList, newRating)
	}

	var blog = model.BlogPage{
		Title:        request.Title,
		Description:  request.Description,
		CreationDate: time.Now(),
		Status:       uint(request.Status),
		UserId:       int(request.UserId),
		RatingSum:    int(request.RatingSum),
		Ratings:      ratingsList,
	}

	if len(blog.Ratings) == 0 {
		blog.Ratings = make([]model.Rating, 0)
	}

	if err := h.DatabaseConnection.Table(`blog."Blogs"`).Create(&blog).Error; err != nil {
		return nil, err
	}

	responseRatings := make([]*blogs.Rating, len(blog.Ratings))
	for i, rating := range blog.Ratings {
		responseRatings[i] = &blogs.Rating{
			UserId:       int32(rating.UserId),
			CreationDate: rating.CreationDate.Format("2006-01-02T15:04:05.000Z"),
			RatingValue:  int32(rating.RatingValue),
		}
	}

	response := &blogs.Blog{
		Id:           int32(blog.Id),
		Title:        blog.Title,
		Description:  blog.Description,
		CreationDate: blog.CreationDate.Format("2006-01-02T15:04:05.000Z"),
		Status:       int32(blog.Status),
		UserId:       int32(blog.UserId),
		RatingSum:    int32(blog.RatingSum),
		Ratings:      responseRatings,
	}

	return response, nil
}

func (h BlogHandler) GetAllBlog(ctx context.Context, request *blogs.Empty) (*blogs.ListBlog, error) {
	var blogsFromDB []model.BlogPage
	if err := h.DatabaseConnection.Table(`blog."Blogs"`).Find(&blogsFromDB).Error; err != nil {
		return nil, err
	}

	// Mapiranje blogova iz baze na proto strukturu
	var protoBlogs []*blogs.Blog
	for _, blog := range blogsFromDB {
		ratingList := make([]*blogs.Rating, len(blog.Ratings))
		for i, rating := range blog.Ratings {
			ratingList[i] = &blogs.Rating{
				UserId:       int32(rating.UserId),
				CreationDate: rating.CreationDate.String(),
				RatingValue:  int32(rating.RatingValue),
			}
		}

		protoBlog := &blogs.Blog{
			Id:           int32(blog.Id),
			Title:        blog.Title,
			Description:  blog.Description,
			CreationDate: blog.CreationDate.String(),
			Status:       int32(blog.Status),
			UserId:       int32(blog.UserId),
			RatingSum:    int32(blog.RatingSum),
			Ratings:      ratingList,
		}

		protoBlogs = append(protoBlogs, protoBlog)
	}

	// Kreiranje odgovora koji sadrži listu blogova
	response := &blogs.ListBlog{
		Blogs: protoBlogs,
	}

	return response, nil
}
