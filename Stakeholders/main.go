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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connStr := "user=postgres dbname=explorer-v1 password=super sslmode=disable"
	//connStr := "user=postgres dbname=explorer password=super sslmode=disable port=5432 host=database"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {

	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

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

	authHandler := handl.AuthHandler{DatabaseConnection: database, Key: "your_secret_key"}
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
