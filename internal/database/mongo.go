package database

import (
	"context"
	"errors"
	"time"

	"crud/config"
	"crud/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
)

func InitMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return err
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	mongoClient = client
	mongoDatabase = client.Database(config.MongoDBName)
	return nil
}

func CloseMongoDB() {
	if mongoClient != nil {
		_ = mongoClient.Disconnect(context.Background())
	}
}

func GetMongoUsers() ([]models.User, error) {
	var users []models.User
	collection := mongoDatabase.Collection("users")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetMongoUserByID(id string) (models.User, error) {
	var user models.User
	collection := mongoDatabase.Collection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, errors.New("invalid user ID format")
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, errors.New("user not found")
	} else if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateMongoUser(user models.User) (string, error) {
	collection := mongoDatabase.Collection("users")

	user.ID = primitive.NewObjectID().Hex() // Assign new ObjectID as string
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return primitive.NilObjectID.Hex(), err
	}

	return user.ID, nil
}

func UpdateMongoUser(id string, updatedUser models.User) error {
	collection := mongoDatabase.Collection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	update := bson.M{
		"$set": bson.M{
			"name":  updatedUser.Name,
			"email": updatedUser.Email,
		},
	}

	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

func DeleteMongoUser(id string) error {
	collection := mongoDatabase.Collection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
