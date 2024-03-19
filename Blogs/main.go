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

func startServer(handler *handler.BlogHandler) {
	router := mux.NewRouter().StrictSlash(true)

	// Dodajemo middleware za CORS
	router.Use(corsMiddleware)

	router.HandleFunc("/blogs", handler.Create).Methods("POST")
	router.HandleFunc("/blogs", handler.GetAllBlogs).Methods("GET")
	router.HandleFunc("/blogs/{id:[+-]?[0-9]+}", handler.GetBlogByID).Methods("GET")
	router.HandleFunc("/blogs/updateOneBlog", handler.Update).Methods("PUT")
	router.HandleFunc("/blogs/getByStatus/{state:[+-]?[0-9]+}", handler.GetAllBlogsByStatus).Methods("GET")
	router.HandleFunc("/blogs/rating/{userId:[+-]?[0-9]+}/{blogId:[+-]?[0-9]+}/{value:[+-]?[0-9]+}", handler.UpdateRating).Methods("PUT")
	router.HandleFunc("/blogs/rating/{userId:[+-]?[0-9]+}/{blogId:[+-]?[0-9]+}", handler.DeleteRating).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Middleware funkcija za CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost:44333/api/")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
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
	handler := &handler.BlogHandler{BlogService: BlogService}
	startServer(handler)

	print("ok")
}
