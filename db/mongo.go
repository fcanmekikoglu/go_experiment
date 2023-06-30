package db

import (
	"context"
	"fmt"
	"time"

	"github.com/fcanmekikoglu/go_experiment/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InsertFact(client *mongo.Client, fact types.Fact, dbName, collectionName string) error {
	collection := client.Database(dbName).Collection(collectionName)

	_, err := collection.InsertOne(context.Background(), bson.M{"fact": fact.Text, "type": fact.Type})
	if err != nil {
		return err
	}
	fmt.Printf("Fact inserted successfully into %s.%s\n", dbName, collectionName)
	return nil
}
