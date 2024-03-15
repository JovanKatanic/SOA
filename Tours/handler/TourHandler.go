package handler

import (
	"encoding/json"
	"net/http"
	"tours_service/model"
	"tours_service/service"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourService.Create(&tour)
	if err != nil {
		println("Error while creating a new facility")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
