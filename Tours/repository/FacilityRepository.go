package repository

import (
	"context"
	"fmt"
	"time"
	"tours_service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FacilityRepository struct {
	FacilityClient *mongo.Client
}

func (rep *FacilityRepository) getCollection() *mongo.Collection {
	facilityDatabase := rep.FacilityClient.Database("mongodb")
	facilitysCollection := facilityDatabase.Collection("facilities")
	return facilitysCollection
}

func (rep *FacilityRepository) Insert(facility *model.Facility) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	facilityCollection := rep.getCollection()

	result, err := facilityCollection.InsertOne(ctx, &facility)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (pr *FacilityRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	facilityCollection := pr.getCollection()

	filter := bson.D{{Key: "_id", Value: id}}
	result, err := facilityCollection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}
