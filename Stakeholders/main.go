package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"stakeholders_service/handl"
	"stakeholders_service/proto/auth"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	authHandler := handl.AuthHandler{}
	auth.RegisterStakeholderServiceServer(grpcServer, authHandler)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
}
