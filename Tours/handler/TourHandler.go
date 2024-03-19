package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tours_service/model"
	"tours_service/service"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) CreateTour(resp http.ResponseWriter, req *http.Request) {
	var tour model.Tour

	err := json.NewDecoder(req.Body).Decode(&tour)

	if err != nil {
		println("Error while parsing json: ", err.Error())
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(tour, tour.Tags)

	err = handler.TourService.CreateTour(&tour)
	if err != nil {
		println("Error while creating a new tour")
		return
	}

	resp.WriteHeader(http.StatusCreated)
	resp.Header().Set("Content-Type", "application/jsons")
}
