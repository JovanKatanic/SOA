package handler

import (
	"blogs_service/model"
	"blogs_service/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (handler *BlogHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var blog model.BlogPage
	err := json.NewDecoder(req.Body).Decode(&blog)
	fmt.Println(len(blog.Ratings))
	fmt.Println(blog.Ratings)
	if len(blog.Ratings) == 0 {
		fmt.Println("usao u if")
		blog.Ratings = []model.Ratings{
			{UserId: 0, CreationDate: time.Time{}, RatingValue: 0},
		}
		fmt.Println(blog.Ratings)
	}
	fmt.Println(blog.Ratings)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.BlogService.Create(&blog)
	if err != nil {
		println("Error while creating a new blog")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
