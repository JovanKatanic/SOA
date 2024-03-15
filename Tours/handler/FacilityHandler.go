package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tours_service/model"
	"tours_service/service"

	"github.com/gorilla/mux"
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

func (handler *FacilityHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])

	if err != nil {
		http.Error(writer, "Invalid facility ID", http.StatusBadRequest)
		return
	}
	err = handler.FacilityService.Delete(id)
	if err != nil {
		println("Error while deleting a facility")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
