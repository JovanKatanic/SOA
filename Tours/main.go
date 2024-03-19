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

func startServer(handler *handler.FacilityHandler, tourHandler *handler.TourHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(corsMiddleware)

	router.HandleFunc("/facilities", handler.Create).Methods("POST")
	router.HandleFunc("/createTour", tourHandler.CreateTour).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func setUpDependencies(database *gorm.DB) (*handler.FacilityHandler, *handler.TourHandler) {
	FacilityRepository := &repository.FacilityRepository{DatabaseConnection: database}
	FacilityService := &service.FacilityService{FacilityRepository: FacilityRepository}
	facilityHandler := &handler.FacilityHandler{FacilityService: FacilityService}

	TourRepository := &repository.TourRepository{DatabaseConnection: database}
	TourService := &service.TourService{TourRepository: TourRepository}
	tourHandler := &handler.TourHandler{TourService: TourService}

	return facilityHandler, tourHandler
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	facilityHandler, tourHandler := setUpDependencies(database)

	startServer(facilityHandler, tourHandler)

	print("ok")
}
