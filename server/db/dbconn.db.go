package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
	DBinstance() (*mongo.Client, bool)
	OpenCollection(client *mongo.Client, database *mongo.Database, collection string) *mongo.Collection
	OpenDatabase(client *mongo.Client, database string) *mongo.Database
}

type db struct{}

func Newdatabase() Database {
	return &db{}
}

func (db *db) DBinstance() (*mongo.Client, bool) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, false

	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Check the connection
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	fmt.Println("Connected to MongoDB!")

	return client, true
}

func (db *db) OpenDatabase(client *mongo.Client, database string) *mongo.Database {
	Database := client.Database(database)
	return Database
}
func (db *db) OpenCollection(client *mongo.Client, database *mongo.Database, collection string) *mongo.Collection {
	Collection := database.Collection(collection)

	return Collection
}
