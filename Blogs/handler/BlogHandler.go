package handler

import (
	"blogs_service/model"
	"blogs_service/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (handler *BlogHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var blog model.BlogPage
	err := json.NewDecoder(req.Body).Decode(&blog)

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
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(blog)
}

func (handler *BlogHandler) GetAllBlogs(writer http.ResponseWriter, req *http.Request) {
	blogs, err := handler.BlogService.GetAllBlogs()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(blogs); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *BlogHandler) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	blog, err := handler.BlogService.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func (handler *BlogHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var blog model.BlogPage
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BlogService.UpdateOneBlog(&blog)
	if err != nil {
		fmt.Println("Error while updating the blog:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
}

func (handler *BlogHandler) GetAllBlogsByStatus(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	status, err := strconv.Atoi(vars["state"])
	if err != nil {
		http.Error(writer, "Invalid blog status", http.StatusBadRequest)
		return
	}

	blogs, err := handler.BlogService.GetAllByStatus(status)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(blogs); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
