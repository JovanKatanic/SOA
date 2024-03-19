package handler

import (
	"blogs_service/model"
	"blogs_service/service"
	"encoding/json"
	"net/http"
)

type CommentHandler struct {
	CommentService *service.CommentService
}

func (handler *CommentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var comment model.Comment
	err := json.NewDecoder(req.Body).Decode(&comment)

	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.CommentService.CreateComment(&comment)
	if err != nil {
		println("Error while creating a new blog")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(comment)
}
