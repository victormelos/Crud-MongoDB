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
	logger.Info("Iniciando conexão com MongoDB")

	mongodb_uri := os.Getenv("MONGODB_URL")
	if mongodb_uri == "" {
		mongodb_uri = "mongodb://localhost:27017"
		logger.Info("Usando URL padrão do MongoDB: " + mongodb_uri)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	logger.Info("Conexão com MongoDB estabelecida com sucesso")
	MongoDBClient = client
	return MongoDBClient, nil
}
