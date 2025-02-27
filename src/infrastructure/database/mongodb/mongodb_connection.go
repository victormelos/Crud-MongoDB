package mongodb

import (
	"context"
	"os"

	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDBConnection cria uma nova conexão com o MongoDB
func NewMongoDBConnection() (*mongo.Client, error) {
	logger.Info("Conectando ao MongoDB")

	mongodbURI := os.Getenv("MONGODB_URL")
	if mongodbURI == "" {
		mongodbURI = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(mongodbURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	logger.Info("Conectado ao MongoDB com sucesso")
	return client, nil
}
