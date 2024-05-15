package main

import (
	"api_gateway/proto/auth"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address                    string
	StakeholdersServiceAddress string
}

func main() {
	cfg := Config{
		StakeholdersServiceAddress: "localhost:8000",
		Address:                    ":8000",
	}

	conn, err := grpc.DialContext(
		context.Background(),
		cfg.StakeholdersServiceAddress,
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                            // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // Allow specific HTTP methods
		AllowedHeaders:   []string{"*"},                            // Allow all headers
		AllowCredentials: true,                                     // Allow sending credentials (e.g., cookies)
	})

	handler := c.Handler(gwmux)

	gwServer := &http.Server{
		Addr:    cfg.Address,
		Handler: handler,
	}

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
