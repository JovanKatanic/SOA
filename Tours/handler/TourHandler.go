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

func (handler *TourHandler) CreateTour(resp http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	var createdTour *model.Tour

	err := json.NewDecoder(req.Body).Decode(&tour)

	if err != nil {
		println("Error while parsing json: ", err.Error())
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	createdTour, err = handler.TourService.CreateTour(&tour)
	if err != nil {
		println("Error while creating a new tour: ", err.Error())
		return
	}

	jsonResponse, err := json.Marshal(createdTour)
	if err != nil {
		println("Error while encoding tour to JSON: ", err.Error())
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusCreated)
	resp.Header().Set("Content-Type", "application/jsons")

	_, err = resp.Write(jsonResponse)
	if err != nil {
		println("Error while writing response: ", err.Error())
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}
