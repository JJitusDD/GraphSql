package repository

import "go.mongodb.org/mongo-driver/mongo"

type MongoRepoImpl struct {
	collectionUsers *mongo.Collection
	collectionTodos *mongo.Collection
}

func NewMongoDB(db *mongo.Client) *MongoRepoImpl {
	districts := db.Database("mongo-data").Collection("users")
	province := db.Database("mongo-data").Collection("provinces")

	return &MongoRepoImpl{
		collectionUsers: districts,
		collectionTodos: province,
	}
}
