package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user_management_service/handler"
	"user_management_service/repository"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8082"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	followerLogger := log.New(os.Stdout, "[follower-api] ", log.LstdFlags)
	followerStoreLogger := log.New(os.Stdout, "[follower-store] ", log.LstdFlags)

	followerStore, err := repository.NewFollowerRepository(followerStoreLogger)
	if err != nil {
		followerLogger.Fatal(err)
	}

	defer followerStore.CloseDriverConnection(timeoutContext)
	followerStore.CheckConnection()

	followerHandler := handler.NewFollowerHandler(followerLogger, followerStore)

	router := mux.NewRouter()
	router.Use(followerHandler.MiddlewareContentTypeSet)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	followerLogger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			followerLogger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	followerLogger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		followerLogger.Fatal("Cannot gracefully shutdown...")
	}
	followerLogger.Println("Server stopped")
}
