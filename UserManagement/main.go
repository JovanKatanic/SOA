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

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	//connStr := "user=postgres dbname=explorer-v1 password=super sslmode=disable"
	connStr := "user=postgres dbname=explorer password=super sslmode=disable port=5432 host=database"
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

	listener, err := net.Listen("tcp", "localhost:8002")
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
	fmt.Println("startovao server")
	reflection.Register(grpcServer)

	blogHandler := handl.FollowingsHandler{DatabaseConnection: database}
	followings.RegisterFollowerServiceServer(grpcServer, blogHandler)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()

	/*BlogRepository := &repository.BlogRepository{DatabaseConnection: database}
	BlogService := &service.BlogService{BlogRepository: BlogRepository}
	BlogHandler := &handler.BlogHandler{BlogService: BlogService}

	CommentRepository := &repository.CommentRepository{DatabaseConnection: database}
	CommentService := &service.CommentService{CommentRepository: CommentRepository}
	CommentHandler := &handler.CommentHandler{CommentService: CommentService}

	startServer(BlogHandler, CommentHandler)

	print("ok")*/
}

// func main() {
// 	port := os.Getenv("PORT")
// 	if len(port) == 0 {
// 		port = "8082"
// 	}

// 	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	followerLogger := log.New(os.Stdout, "[follower-api] ", log.LstdFlags)
// 	followerStoreLogger := log.New(os.Stdout, "[follower-store] ", log.LstdFlags)

// 	followerStore, err := repository.NewFollowerRepository(followerStoreLogger)
// 	if err != nil {
// 		followerLogger.Fatal(err)
// 	}

// 	defer followerStore.CloseDriverConnection(timeoutContext)
// 	followerStore.CheckConnection()

// 	followerHandler := handler.NewFollowerHandler(followerLogger, followerStore)

// 	router := mux.NewRouter()
// 	router.Use(followerHandler.MiddlewareContentTypeSet)

// 	postFollowRelationship := router.Methods(http.MethodPost).Subrouter()
// 	postFollowRelationship.HandleFunc("/createFollow", followerHandler.CreateFollow)
// 	postFollowRelationship.Use(followerHandler.MiddlewareFollowerDeserialization)

// 	getFollowingss := router.Methods(http.MethodGet).Subrouter()
// 	getFollowingss.HandleFunc("/followings/{id}", followerHandler.GetAllFollowings)

// 	getRecommendedFollowings := router.Methods(http.MethodGet).Subrouter()
// 	getRecommendedFollowings.HandleFunc("/recommendedfollowings/{id}", followerHandler.GetAllRecommendedFollowings)

// 	deleteFollowRelationship := router.Methods(http.MethodDelete).Subrouter()
// 	deleteFollowRelationship.HandleFunc("/deleteFollow/{followerId}/{followedId}", followerHandler.DeleteFollow)

// 	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

// 	server := http.Server{
// 		Addr:         ":" + port,
// 		Handler:      cors(router),
// 		IdleTimeout:  120 * time.Second,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 5 * time.Second,
// 	}

// 	followerLogger.Println("Server listening on port", port)

// 	go func() {
// 		err := server.ListenAndServe()
// 		if err != nil {
// 			followerLogger.Fatal(err)
// 		}
// 	}()

// 	sigCh := make(chan os.Signal)
// 	signal.Notify(sigCh, os.Interrupt)
// 	signal.Notify(sigCh, os.Kill)

// 	sig := <-sigCh
// 	followerLogger.Println("Received terminate, graceful shutdown", sig)

// 	if server.Shutdown(timeoutContext) != nil {
// 		followerLogger.Fatal("Cannot gracefully shutdown...")
// 	}
// 	followerLogger.Println("Server stopped")
// }
