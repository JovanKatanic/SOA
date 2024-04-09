package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"tours_service/handler"
	"tours_service/repository"
	"tours_service/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	gorillaHandlers "github.com/gorilla/handlers"
)

func initMongoDb() *mongo.Client {

	dburi := "mongodb://mongo:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		fmt.Print(err)
	}

	return client
}
func manageRouter(client *mongo.Client) http.Server {
	FacilityRepository := &repository.FacilityRepository{FacilityClient: client}
	FacilityService := &service.FacilityService{FacilityRepository: FacilityRepository}
	FacilityHandler := &handler.FacilityHandler{FacilityService: FacilityService}

	KeypointRepository := &repository.KeypointRepository{KeypointClient: client}
	KeypointService := &service.KeypointService{KeypointRepository: KeypointRepository}
	KeypointHandler := &handler.KeypointHandler{KeypointService: KeypointService}

	tourRepository := &repository.TourRepository{TourClient: client}
	tourService := &service.TourService{TourRepository: tourRepository}
	tourHandler := &handler.TourHandler{TourService: tourService}

	tourRatingRepository := &repository.TourRatingRepository{TourRatingClient: client}
	tourRatingService := &service.TourRatingService{TourRatingRepository: tourRatingRepository}
	tourRatingHandler := &handler.TourRatingHandler{TourRatingService: tourRatingService}

	tourProblemRepository := &repository.TourProblemRepository{TourProblemClient: client}
	tourProblemService := &service.TourProblemService{TourProblemRepository: tourProblemRepository,
		TourService: tourService}
	tourProblemHandler := &handler.TourProblemHandler{TourProblemService: tourProblemService}

	router := mux.NewRouter().StrictSlash(true)

	postFacilityRouter := router.Methods(http.MethodPost).Subrouter()
	postFacilityRouter.HandleFunc("/facilities", FacilityHandler.CreateFacility)
	postFacilityRouter.Use(FacilityHandler.MiddlewareFacilityDeserialization)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/facilities/{id}", FacilityHandler.DeleteFacility)

	postKeypointRouter := router.Methods(http.MethodPost).Subrouter()
	postKeypointRouter.HandleFunc("/keypoints", KeypointHandler.CreateKeypoint)
	postKeypointRouter.Use(KeypointHandler.MiddlewareKeypointDeserialization)

	postTourRouter := router.Methods(http.MethodPost).Subrouter()
	postTourRouter.HandleFunc("/createTour", tourHandler.CreateTour)
	postTourRouter.Use(tourHandler.MiddlewareTourDeserialization)

	putTourRouter := router.Methods(http.MethodPut).Subrouter()
	putTourRouter.HandleFunc("/tours", tourHandler.UpdateTour)
	putTourRouter.Use(tourHandler.MiddlewareTourDeserialization)

	getTourRouter := router.Methods(http.MethodGet).Subrouter()
	getTourRouter.HandleFunc("/tours/{id}", tourHandler.GetTourById)

	postTourRatingRouter := router.Methods(http.MethodPost).Subrouter()
	postTourRatingRouter.HandleFunc("/createTourRating", tourRatingHandler.CreateTourRating)
	postTourRatingRouter.Use(tourRatingHandler.MiddlewareTourRatingDeserialization)

	getTourProblemRouter := router.Methods(http.MethodGet).Subrouter()
	getTourProblemRouter.HandleFunc("/getByAuthorId/{authorId}", tourProblemHandler.GetByAuthorId)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":8080",
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	return server

}

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := initMongoDb()

	err := client.Connect(timeoutContext)
	if err != nil {
		fmt.Print(err)
	}
	logger := log.New(os.Stdout, "[logger-main] ", log.LstdFlags)
	server := manageRouter(client)

	logger.Println("Server listening on port", "8080")
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")

	print("ok")
}
