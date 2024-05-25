package main

import (
	"api_gateway/proto/auth"
	"api_gateway/proto/blogs"
	"api_gateway/proto/followings"
	"api_gateway/proto/tours"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	StakeholderServiceAddress string
	Address                   string
	BlogServiceAddress        string
	BlogAddress               string
	FollowingServiceAdress    string
	FOllowingAddress          string
	ToursServiceAddress       string
	TourAddress               string
}

func main() {
	cfg := Config{
		BlogServiceAddress:        "localhost:8001",
		BlogAddress:               ":8001",
		FollowingServiceAdress:    "localhost:8002",
		FOllowingAddress:          ":8002",
		StakeholderServiceAddress: "localhost:8000",
		Address:                   ":8000",
		ToursServiceAddress:       "localhost:8003",
		TourAddress:               ":8003",
	}

	conn, err := grpc.DialContext(
		context.Background(),
		cfg.StakeholderServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	client := auth.NewStakeholderServiceClient(conn)
	err = auth.RegisterStakeholderServiceHandlerClient(
		context.Background(),
		gwmux,
		client,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    cfg.Address,
		Handler: gwmux,
	}

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	blogConn, err := grpc.DialContext(
		context.Background(),
		cfg.BlogServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial blog service:", err)
	}

	blogGwmux := runtime.NewServeMux()
	blogClient := blogs.NewBlogServiceClient(blogConn)
	err = blogs.RegisterBlogServiceHandlerClient(
		context.Background(),
		blogGwmux,
		blogClient,
	)
	if err != nil {
		log.Fatalln("Failed to register blog gateway:", err)
	}

	blogGwServer := &http.Server{
		Addr:    cfg.BlogAddress,
		Handler: blogGwmux,
	}

	go func() {
		if err := blogGwServer.ListenAndServe(); err != nil {
			log.Fatal("Blog server error: ", err)
		}
	}()

	followCon, err := grpc.DialContext(
		context.Background(),
		cfg.FollowingServiceAdress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial blog service:", err)
	}

	followGMux := runtime.NewServeMux()
	followClient := followings.NewFollowerServiceClient(followCon)
	err = followings.RegisterFollowerServiceHandlerClient(
		context.Background(),
		followGMux,
		followClient,
	)
	if err != nil {
		log.Fatalln("Failed to register blog gateway:", err)
	}

	followingGwServer := &http.Server{
		Addr:    cfg.FOllowingAddress,
		Handler: followGMux,
	}

	go func() {
		if err := followingGwServer.ListenAndServe(); err != nil {
			log.Fatal("Blog server error: ", err)
		}
	}()

	tourConn, err := grpc.DialContext(
		context.Background(),
		cfg.ToursServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial tour service:", err)
	}

	tlourGMux := runtime.NewServeMux()
	tourClient := tours.NewTourServiceClient(tourConn)
	err = tours.RegisterTourServiceHandlerClient(
		context.Background(),
		tlourGMux,
		tourClient,
	)
	if err != nil {
		log.Fatalln("Failed to register tourr gateway:", err)
	}

	tourGwServer := &http.Server{
		Addr:    cfg.TourAddress,
		Handler: tlourGMux,
	}

	go func() {
		if err := tourGwServer.ListenAndServe(); err != nil {
			log.Fatal("Tour server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
