package database

import (
	"context"
	"errors"

	"crud/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB(client *mongo.Client, dbName string) *MongoDB {
	return &MongoDB{
		client:   client,
		database: client.Database(dbName),
	}
}

func (db *MongoDB) GetUsers() ([]models.User, error) {
	var users []models.User
	collection := db.database.Collection("users")

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

func (db *MongoDB) GetUserByID(id string) (models.User, error) {
	var user models.User
	collection := db.database.Collection("users")

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

func (db *MongoDB) CreateUser(user models.User) (string, error) {
	collection := db.database.Collection("users")

	user.ID = primitive.NewObjectID().Hex()
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (db *MongoDB) UpdateUser(id string, updatedUser models.User) error {
	collection := db.database.Collection("users")

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

func (db *MongoDB) DeleteUser(id string) error {
	collection := db.database.Collection("users")

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
