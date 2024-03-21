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
	connStr := "user=postgres dbname=explorer password=super sslmode=disable port=5432 host=database"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func startServer(FacilityHandler *handler.FacilityHandler, KeypointHandler *handler.KeypointHandler, tourHandler *handler.TourHandler, tourRatingHandler *handler.TourRatingHandler, tourProblemHandler *handler.TourProblemHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/facilities", FacilityHandler.Create).Methods("POST")
	router.HandleFunc("/facilities/{id}", FacilityHandler.Delete).Methods("DELETE")

	router.HandleFunc("/keypoints", KeypointHandler.Create).Methods("POST")

	router.HandleFunc("/createTour", tourHandler.CreateTour).Methods("POST")
	router.HandleFunc("/tours", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/tours/{id}", tourHandler.GetTourByID).Methods("GET")

	router.HandleFunc("/createTourRating", tourRatingHandler.CreateTourRating).Methods("POST")

	router.HandleFunc("/getByAuthorId/{authorId}", tourProblemHandler.GetByAuthorId).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
	println("Server starting")

	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(router)))
}

func setUpDependencies(database *gorm.DB) (*handler.FacilityHandler, *handler.TourHandler, *handler.KeypointHandler, *handler.TourRatingHandler, *handler.TourProblemHandler) {
	FacilityRepository := &repository.FacilityRepository{DatabaseConnection: database}
	FacilityService := &service.FacilityService{FacilityRepository: FacilityRepository}
	FacilityHandler := &handler.FacilityHandler{FacilityService: FacilityService}

	KeypointRepository := &repository.KeypointRepository{DatabaseConnection: database}
	KeypointService := &service.KeypointService{KeypointRepository: KeypointRepository}
	KeypointHandler := &handler.KeypointHandler{KeypointService: KeypointService}

	tourRepository := &repository.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepository: tourRepository}
	tourHandler := &handler.TourHandler{TourService: tourService}

	tourRatingRepository := &repository.TourRatingRepository{DatabaseConnection: database}
	tourRatingService := &service.TourRatingService{TourRatingRepository: tourRatingRepository}
	tourRatingHandler := &handler.TourRatingHandler{TourRatingService: tourRatingService}

	tourProblemRepository := &repository.TourProblemRepository{DatabaseConnection: database}
	tourProblemService := &service.TourProblemService{TourProblemRepository: tourProblemRepository,
		TourService: tourService}
	tourProblemHandler := &handler.TourProblemHandler{TourProblemService: tourProblemService}

	return FacilityHandler, tourHandler, KeypointHandler, tourRatingHandler, tourProblemHandler
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	facilityHandler, tourHandler, keypointHandler, tourRatingHandler, tourProblemHandler := setUpDependencies(database)

	startServer(facilityHandler, keypointHandler, tourHandler, tourRatingHandler, tourProblemHandler)

	print("ok")
}
