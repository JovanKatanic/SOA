package handler

import (
	"blogs_service/model"
	"blogs_service/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type CommentHandler struct {
	CommentService *service.CommentService
}

func (handler *CommentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var comment model.Comment
	err := json.NewDecoder(req.Body).Decode(&comment)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.CommentService.CreateComment(&comment)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(comment)
}

func (handler *CommentHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var comment model.Comment
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CommentService.UpdateComment(&comment)
	if err != nil {
		fmt.Println("Error while updating the blog:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(comment)
}
