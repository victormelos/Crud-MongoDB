package mongodb

import (
	"context"
	"os"

	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDBClient *mongo.Client
)

func NewMongoDBConnection() (*mongo.Client, error) {
	mongodb_uri := os.Getenv("MONGODB_URL")
	if mongodb_uri == "" {
		mongodb_uri = ""

	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(context.Background()); err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	logger.Info("Database connected")
	MongoDBClient = client
	return MongoDBClient, nil
}
