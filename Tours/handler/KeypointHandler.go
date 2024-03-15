package handler

import (
	"encoding/json"
	"net/http"
	"tours_service/model"
	"tours_service/service"
)

type KeypointHandler struct {
	KeypointService *service.KeypointService
}

func (handler *KeypointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var keypoint model.Keypoint
	err := json.NewDecoder(req.Body).Decode(&keypoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.KeypointService.Create(&keypoint)
	if err != nil {
		println("Error while creating a new facility")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
