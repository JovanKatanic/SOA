package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user_management_service/handl"
	"user_management_service/proto/followings"
	"user_management_service/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()

	followerLogger := log.New(os.Stdout, "[follower-api] ", log.LstdFlags)
	followerStoreLogger := log.New(os.Stdout, "[follower-store] ", log.LstdFlags)

	// Inicijalizacija Neo4j baze
	/*followerStore, err := repository.NewFollowerRepository(followerStoreLogger)
	if err != nil {
		followerLogger.Fatal(err)
	}
	defer followerStore.CloseDriverConnection(timeoutContext)*/
	uri := "bolt://localhost:7687" // Promenite na odgovarajuću adresu vaše baze
	user := "neo4j"                // Promenite na odgovarajuće korisničko ime
	pass := "sifra123"             // Promenite na odgovarajuću lozinku

	followerStore, err := repository.NewFollowerRepository(followerStoreLogger, uri, user, pass)
	if err != nil {
		followerLogger.Fatal(err)
	}
	followerStore.CheckConnection()

	/*followerHandler := handler.NewFollowerHandler(followerLogger, followerStore)

	router := mux.NewRouter()
	router.Use(followerHandler.MiddlewareContentTypeSet)

	postFollowRelationship := router.Methods(http.MethodPost).Subrouter()
	postFollowRelationship.HandleFunc("/createFollow", followerHandler.CreateFollow)
	postFollowRelationship.Use(followerHandler.MiddlewareFollowerDeserialization)

	getFollowingss := router.Methods(http.MethodGet).Subrouter()
	getFollowingss.HandleFunc("/followings/{id}", followerHandler.GetAllFollowings)

	getRecommendedFollowings := router.Methods(http.MethodGet).Subrouter()
	getRecommendedFollowings.HandleFunc("/recommendedfollowings/{id}", followerHandler.GetAllRecommendedFollowings)

	deleteFollowRelationship := router.Methods(http.MethodDelete).Subrouter()
	deleteFollowRelationship.HandleFunc("/deleteFollow/{followerId}/{followedId}", followerHandler.DeleteFollow)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))*/

	// gRPC Server setup
	grpcListener, err := net.Listen("tcp", "localhost:8003")
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcListener.Close()

	grpcServer := grpc.NewServer()
	fmt.Println("UserManagement server started")
	reflection.Register(grpcServer)

	/*followerHandler := handler.NewFollowerHandler(followerLogger, followerStore)

	router := mux.NewRouter()
	router.Use(followerHandler.MiddlewareContentTypeSet)

	postFollowRelationship := router.Methods(http.MethodPost).Subrouter()
	postFollowRelationship.HandleFunc("/createFollow", followerHandler.CreateFollow)
	postFollowRelationship.Use(followerHandler.MiddlewareFollowerDeserialization)*/

	followingsHandler := handl.NewFollowingsHandler(followerLogger, followerStore)
	followings.RegisterFollowerServiceServer(grpcServer, followingsHandler)

	//gRPC servers in separate goroutines
	go func() {
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatal("gRPC server error: ", err)
		}
	}()

	// Wait for termination signal
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	sig := <-sigCh
	followerLogger.Println("Received terminate signal, graceful shutdown", sig)

	// Graceful shutdown for gRPC server
	grpcServer.GracefulStop()
	followerLogger.Println("gRPC server stopped")
}
