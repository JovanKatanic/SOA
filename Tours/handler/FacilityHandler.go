package handler

import (
	"encoding/json"
	"net/http"
	"tours_service/model"
	"tours_service/service"
)

type FacilityHandler struct {
	FacilityService *service.FacilityService
}

func (handler *FacilityHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var facility model.Facility
	err := json.NewDecoder(req.Body).Decode(&facility)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.FacilityService.Create(&facility)
	if err != nil {
		println("Error while creating a new facility")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
