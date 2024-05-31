package main

import (
	"api_gateway/proto/auth"
	"api_gateway/proto/blogs"
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
}

func main() {
	cfg := Config{
		StakeholderServiceAddress: "localhost:8000",
		Address:                   ":8000",
		BlogServiceAddress:        "localhost:8001",
	}

	gwmux := runtime.NewServeMux()

	// Connect to the Stakeholder Service
	conn, err := grpc.DialContext(
		context.Background(),
		cfg.StakeholderServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial stakeholder server:", err)
	}
	defer conn.Close()

	stakeholderClient := auth.NewStakeholderServiceClient(conn)
	err = auth.RegisterStakeholderServiceHandlerClient(
		context.Background(),
		gwmux,
		stakeholderClient,
	)
	if err != nil {
		log.Fatalln("Failed to register stakeholder gateway:", err)
	}

	// Connect to the Blog Service
	blogConn, err := grpc.DialContext(
		context.Background(),
		cfg.BlogServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial blog server:", err)
	}
	defer blogConn.Close()

	blogClient := blogs.NewBlogServiceClient(blogConn)
	err = blogs.RegisterBlogServiceHandlerClient(
		context.Background(),
		gwmux,
		blogClient,
	)
	if err != nil {
		log.Fatalln("Failed to register blog gateway:", err)
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

	// Graceful shutdown
	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
