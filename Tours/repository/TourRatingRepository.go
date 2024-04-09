package repository

import (
	"context"
	"fmt"
	"time"
	"tours_service/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type TourRatingRepository struct {
	TourRatingClient *mongo.Client
}

func (rep *TourRatingRepository) getCollection() *mongo.Collection {
	tourRatingDatabase := rep.TourRatingClient.Database("mongodb")
	tourRatingCollection := tourRatingDatabase.Collection("tourRatings")
	return tourRatingCollection
}

// func (repo *TourRatingRepository) CreateTourRating(tourRating *model.TourRating) error {
// 	dbResult := repo.DatabaseConnection.Table(`tours."TourRatings"`).Create(tourRating)

// 	if dbResult.Error != nil {
// 		return dbResult.Error
// 	}

// 	println("Rows affected: ", dbResult.RowsAffected)
// 	return nil
// }

func (rep *TourRatingRepository) Insert(tourRating *model.TourRating) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tourRatingCollection := rep.getCollection()

	result, err := tourRatingCollection.InsertOne(ctx, &tourRating)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}
