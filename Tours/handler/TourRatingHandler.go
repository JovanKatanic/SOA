package handler

import (
	"encoding/json"
	"net/http"
	"tours_service/model"
	"tours_service/service"
)

type TourRatingHandler struct {
	TourRatingService *service.TourRatingService
}

func (handler *TourRatingHandler) CreateTourRating(resp http.ResponseWriter, req *http.Request) {
	var tourRating model.TourRating

	err := json.NewDecoder(req.Body).Decode(&tourRating)
	if err != nil {
		println("Error while parsing json: ", err.Error())
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourRatingService.CreateTourRating(&tourRating)
	if err != nil {
		println("Error while creating a new tour rating: ", err.Error())
		return
	}

	resp.WriteHeader(http.StatusCreated)
	resp.Header().Set("Content-Type", "application/jsons")
}
