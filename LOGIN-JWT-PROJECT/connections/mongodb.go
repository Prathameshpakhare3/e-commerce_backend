package connections

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Atlassian01:A_user123@cluster0.32nif.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOptions) // Use mongo.Connect instead of mongo.NewClient
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	// Ping the database to ensure the connection is successful
	if err = Client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("Failed to ping MongoDB: %v", err))
	}

	fmt.Println("Connected to MongoDB!")
}

func GetCollection(database, collection string) *mongo.Collection {
	return Client.Database(database).Collection(collection)
}
