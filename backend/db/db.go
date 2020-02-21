package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetMongo is used to create the conneciton to Mongo
func GetMongo(col string) *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	collection := client.Database("eks").Collection(col)
	if err != nil {
		panic(err)
	}
	return collection
}
