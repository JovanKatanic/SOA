package main

import (
	"log"
	"net/http"
	"tours_service/handler"
	"tours_service/repository"
	"tours_service/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connStr := "user=postgres dbname=explorer-v1 password=super sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func startServer(FacilityHandler *handler.FacilityHandler, KeypointHandler *handler.KeypointHandler, TourHandler *handler.TourHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/facilities", FacilityHandler.Create).Methods("POST")
	router.HandleFunc("/facilities/{id}", FacilityHandler.Delete).Methods("DELETE")

	router.HandleFunc("/keypoints", KeypointHandler.Create).Methods("POST")

	router.HandleFunc("/tours", TourHandler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	FacilityRepository := &repository.FacilityRepository{DatabaseConnection: database}
	FacilityService := &service.FacilityService{FacilityRepository: FacilityRepository}
	FacilityHandler := &handler.FacilityHandler{FacilityService: FacilityService}

	KeypointRepository := &repository.KeypointRepository{DatabaseConnection: database}
	KeypointService := &service.KeypointService{KeypointRepository: KeypointRepository}
	KeypointHandler := &handler.KeypointHandler{KeypointService: KeypointService}

	TourRepository := &repository.TourRepository{DatabaseConnection: database}
	TourService := &service.TourService{TourRepository: TourRepository}
	TourHandler := &handler.TourHandler{TourService: TourService}

	startServer(FacilityHandler, KeypointHandler, TourHandler)

	print("ok")
}
