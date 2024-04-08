package repository

import (
	"context"
	"fmt"
	"time"
	"tours_service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type FacilityRepository struct {
	FacilityClient *mongo.Client
}

func (rep *FacilityRepository) getCollection() *mongo.Collection {
	facilityDatabase := rep.FacilityClient.Database("mongodb")
	facilitysCollection := facilityDatabase.Collection("facilities")
	return facilitysCollection
}

func (pr *FacilityRepository) Disconnect(ctx context.Context) error {
	err := pr.FacilityClient.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (pr *FacilityRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := pr.FacilityClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Print(err)
	}

	databases, err := pr.FacilityClient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(databases)
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

// func (repo *FacilityRepository) CreateFacility(facility *model.Facility) error {
// 	dbResult := repo.DatabaseConnection.Table(`tours."Facilities"`).Create(facility)

// 	if dbResult.Error != nil {
// 		return dbResult.Error
// 	}
// 	println("Rows affected: ", dbResult.RowsAffected)
// 	return nil
// }

// func (repo *FacilityRepository) DeleteFacility(facilityId int) error {

// 	dbResult := repo.DatabaseConnection.Table(`tours."Facilities"`).Delete(&model.Facility{}, facilityId)

// 	if dbResult.Error != nil {
// 		return dbResult.Error
// 	}
// 	println("Rows affected: ", dbResult.RowsAffected)
// 	return nil
// }
