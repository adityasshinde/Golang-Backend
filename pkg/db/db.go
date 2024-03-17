package db

import (
	"context"
	"log"

	"github.com/adityasshinde/Golang-Backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://adityashindebusiness:00000000@beta-dev.clibhdo.mongodb.net/?retryWrites=true&w=majority")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllCitizens() ([]models.Citizen, error) {
	var citizens []models.Citizen
	collection := client.Database("citizensDB").Collection("citizens")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var citizen models.Citizen
		err := cur.Decode(&citizen)
		if err != nil {
			return nil, err
		}
		citizens = append(citizens, citizen)
	}
	return citizens, nil
}

func GetCitizenByID(id string) (models.Citizen, error) {
	var citizen models.Citizen
	collection := client.Database("citizensDB").Collection("citizens")
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&citizen)
	if err != nil {
		return models.Citizen{}, err
	}
	return citizen, nil
}

func CreateCitizen(citizen models.Citizen) error {
	collection := client.Database("citizensDB").Collection("citizens")
	_, err := collection.InsertOne(context.Background(), citizen)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCitizen(id string, citizen models.Citizen) error {
	collection := client.Database("citizensDB").Collection("citizens")
	_, err := collection.ReplaceOne(context.Background(), bson.M{"id": id}, citizen)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCitizen(id string) error {
	collection := client.Database("citizensDB").Collection("citizens")
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}
