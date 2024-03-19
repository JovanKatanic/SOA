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

	router.HandleFunc("/createTour", TourHandler.CreateTour).Methods("POST")
	router.HandleFunc("/tours", TourHandler.Update).Methods("PUT")

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

func setUpDependencies(database *gorm.DB) (*handler.FacilityHandler, *handler.TourHandler, *handler.KeypointHandler) {
	FacilityRepository := &repository.FacilityRepository{DatabaseConnection: database}
	FacilityService := &service.FacilityService{FacilityRepository: FacilityRepository}
	FacilityHandler := &handler.FacilityHandler{FacilityService: FacilityService}

	KeypointRepository := &repository.KeypointRepository{DatabaseConnection: database}
	KeypointService := &service.KeypointService{KeypointRepository: KeypointRepository}
	KeypointHandler := &handler.KeypointHandler{KeypointService: KeypointService}

	TourRepository := &repository.TourRepository{DatabaseConnection: database}
	TourService := &service.TourService{TourRepository: TourRepository}
	TourHandler := &handler.TourHandler{TourService: TourService}

	return FacilityHandler, TourHandler, KeypointHandler
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	FacilityHandler, TourHandler, KeypointHandler := setUpDependencies(database)

	startServer(FacilityHandler, KeypointHandler, TourHandler)

	print("ok")
}
