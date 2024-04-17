package handler

import (
	"log"
	"net/http"
	"user_management_service/repository"
)

type FollowerHandler struct {
	logger *log.Logger
	repo   *repository.FollowerRepository
}

func NewFollowerHandler(log *log.Logger, repo *repository.FollowerRepository) *FollowerHandler {
	return &FollowerHandler{log, repo}
}

func (m *FollowerHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		m.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
