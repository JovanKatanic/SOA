package main

import (
	"stakeholders_service/model"
	"stakeholders_service/repository"
	"stakeholders_service/service"

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

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	UserRepository := &repository.UserRepostiroy{DatabaseConnection: database}
	PersonRepository := &repository.PersonRepository{DatabaseConnection: database}
	AuthenticationService := &service.AuthenticationService{UserRepository: UserRepository, PersonRepository: PersonRepository}
	authen := model.AccountRegistration{
		Username: "lukastankovic",
		Password: "jovan123",
		Email:    "jovan@gmail.com",
		Name:     "Jovan",
		Surname:  "Sarac",
	}
	AuthenticationService.RegisterTourist(&authen)

	print("ok")
}
