package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tours_service/model"
	"tours_service/service"

	"github.com/gorilla/mux"
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

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {

	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// if tour.ID == 0 {
	// 	err = handler.TourService.Create(&tour)
	// 	if err != nil {
	// 		println("Error while creating a new tour")
	// 		writer.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	writer.WriteHeader(http.StatusCreated)
	// } else {
	err = handler.TourService.Update(&tour)
	if err != nil {
		println("Error while updating the tour")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	//}
	responseBody, err := json.Marshal(tour)
	if err != nil {
		println("Error while marshaling tour object to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseBody)
}

func (handler *TourHandler) GetTourByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}

	tour, err := handler.TourService.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tour)
}
