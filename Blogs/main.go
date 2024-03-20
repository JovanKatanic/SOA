package main

import (
	"blogs_service/handler"
	"blogs_service/repository"
	"blogs_service/service"
	"log"
	"net/http"

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

func startServer(BlogHandler *handler.BlogHandler, CommentHandler *handler.CommentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(corsMiddleware)

	router.HandleFunc("/blogs", BlogHandler.Create).Methods("POST")
	router.HandleFunc("/blogs", BlogHandler.GetAllBlogs).Methods("GET")
	router.HandleFunc("/blogs/{id:[+-]?[0-9]+}", BlogHandler.GetBlogByID).Methods("GET")
	router.HandleFunc("/blogs/updateOneBlog", BlogHandler.Update).Methods("PUT")
	router.HandleFunc("/blogs/getByStatus/{state:[+-]?[0-9]+}", BlogHandler.GetAllBlogsByStatus).Methods("GET")
	router.HandleFunc("/blogs/rating/{userId:[+-]?[0-9]+}/{blogId:[+-]?[0-9]+}/{value:[+-]?[0-9]+}", BlogHandler.UpdateRating).Methods("PUT")
	router.HandleFunc("/blogs/rating/{userId:[+-]?[0-9]+}/{blogId:[+-]?[0-9]+}", BlogHandler.DeleteRating).Methods("DELETE")

	router.HandleFunc("/comment", CommentHandler.Create).Methods("POST")
	router.HandleFunc("/comment", CommentHandler.Update).Methods("PUT")
	router.HandleFunc("/comment/{id}", CommentHandler.Delete).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost:44333/api/")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	BlogRepository := &repository.BlogRepository{DatabaseConnection: database}
	BlogService := &service.BlogService{BlogRepository: BlogRepository}
	BlogHandler := &handler.BlogHandler{BlogService: BlogService}

	CommentRepository := &repository.CommentRepository{DatabaseConnection: database}
	CommentService := &service.CommentService{CommentRepository: CommentRepository}
	CommentHandler := &handler.CommentHandler{CommentService: CommentService}

	startServer(BlogHandler, CommentHandler)

	print("ok")
}
