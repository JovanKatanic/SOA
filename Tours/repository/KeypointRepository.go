package repository

import (
	"context"
	"fmt"
	"time"
	"tours_service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type KeypointRepository struct {
	KeypointClient *mongo.Client
}

func (rep *KeypointRepository) getCollection() *mongo.Collection {
	keypointDatabase := rep.KeypointClient.Database("mongodb")
	keypointssCollection := keypointDatabase.Collection("keypoints")
	return keypointssCollection
}

func (rep *KeypointRepository) Insert(keypoint *model.Keypoint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	keypointCollection := rep.getCollection()

	result, err := keypointCollection.InsertOne(ctx, &keypoint)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}
func (pr *KeypointRepository) GetByTourId(id int) ([]model.Keypoint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	keypointCollection := pr.getCollection()

	var keypoints []model.Keypoint
	result, err := keypointCollection.Find(ctx, bson.M{"tourId": id})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err := result.All(ctx, &keypoints); err != nil {
		// Handle decoding error
		fmt.Println(err)
	}

	return keypoints, nil
}
