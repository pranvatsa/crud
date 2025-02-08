package database

import (
	"context"
	"crud/config"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("called here")
	if config.MongoURI == "" {
		return errors.New("❌ MONGO_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	log.Println("✅ Connected to MongoDB")
	DB = client
	return nil
}

func GetCollection(name string) (*mongo.Collection, error) {
	if DB == nil {
		return nil, errors.New("❌ Database connection is not initialized")
	}
	return DB.Database("crudDB").Collection(name), nil
}
