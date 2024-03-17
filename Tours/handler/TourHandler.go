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

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {

	// body, err4 := ioutil.ReadAll(req.Body)
	// fmt.Println("Raw JSON:", string(body))
	// print("----")
	// print(err4)

	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if tour.ID == 0 {
		err = handler.TourService.Create(&tour)
		if err != nil {
			println("Error while creating a new tour")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	} else {
		err = handler.TourService.Update(&tour)
		if err != nil {
			println("Error while updating the tour")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	writer.Header().Set("Content-Type", "application/json")
}
