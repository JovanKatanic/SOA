package handler

import (
	"context"
	"fmt"
	"net/http"
	"tours_service/model"
	"tours_service/service"
)

type KeypointHandler struct {
	KeypointService *service.KeypointService
}

func (handler *KeypointHandler) CreateKeypoint(writer http.ResponseWriter, req *http.Request) {
	keypointInterface := req.Context().Value(KeyProduct{})
	if keypointInterface == nil {
		http.Error(writer, "Keypoint not found in context", http.StatusInternalServerError)
		return
	}
	keypoint, ok := keypointInterface.(*model.Keypoint)
	if !ok {
		http.Error(writer, "Invalid keypoint type in context", http.StatusInternalServerError)
		return
	}
	handler.KeypointService.Create(keypoint)
	writer.WriteHeader(http.StatusCreated)
}

func (p *KeypointHandler) MiddlewareKeypointDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		keypoint := &model.Keypoint{}
		err := keypoint.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			fmt.Print(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, keypoint)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
