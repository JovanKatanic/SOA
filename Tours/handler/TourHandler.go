package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tours_service/model"
	"tours_service/service"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) CreateTour(writer http.ResponseWriter, req *http.Request) {
	tourInterface := req.Context().Value(KeyProduct{})
	if tourInterface == nil {
		http.Error(writer, "Tour not found in context", http.StatusInternalServerError)
		return
	}
	tour, ok := tourInterface.(*model.Tour)
	if !ok {
		http.Error(writer, "Invalid tour type in context", http.StatusInternalServerError)
		return
	}
	handler.TourService.CreateTour(tour)
	tourJSON, err := json.Marshal(tour)
	if err != nil {

		http.Error(writer, "Failed to marshal tour to JSON", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	_, err = writer.Write(tourJSON)
	if err != nil {

		http.Error(writer, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (p *TourHandler) GetTourById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	idstr := vars["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	tour, err := p.TourService.GetTourById(id)
	if err != nil {
		fmt.Print("Database exception: ", err)
	}

	if tour == nil {
		http.Error(rw, "Tour with given id not found", http.StatusNotFound)
		fmt.Printf("Tour with id: '%d' not found", id)
		return
	}

	err = tour.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		fmt.Print("Unable to convert to json :", err)
		return
	}
}
func (handler *TourHandler) UpdateTour(writer http.ResponseWriter, req *http.Request) {
	tourInterface := req.Context().Value(KeyProduct{})
	if tourInterface == nil {
		http.Error(writer, "Tour not found in context", http.StatusInternalServerError)
		return
	}
	tour, ok := tourInterface.(*model.Tour)
	if !ok {
		http.Error(writer, "Invalid tour type in context", http.StatusInternalServerError)
		return
	}
	handler.TourService.UpdateTour(tour)
	writer.WriteHeader(http.StatusCreated)
}

func (p *TourHandler) MiddlewareTourDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		tour := &model.Tour{}
		err := tour.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			fmt.Print(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, tour)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
