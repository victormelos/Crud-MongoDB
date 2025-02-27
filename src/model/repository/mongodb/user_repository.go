package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseConnection *mongo.Client
}

func NewUserRepository(database *mongo.Client) user.UserRepositoryInterface {
	return &userRepository{
		databaseConnection: database,
	}
}

func (ur *userRepository) Create(userDomain *user.UserDomain) (*user.UserDomain, *rest_err.RestErr) {
	collection := ur.databaseConnection.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	userDomain.CreatedAt = time.Now()
	userDomain.UpdatedAt = time.Now()

	logger.Info("Attempting to insert user into MongoDB")
	_, err := collection.InsertOne(context.Background(), userDomain)
	if err != nil {
		logger.Error("Error inserting user into MongoDB", err)
		return nil, rest_err.NewInternalServerError("Error trying to create user")
	}

	logger.Info("User successfully inserted into MongoDB")
	return userDomain, nil
}

func (ur *userRepository) FindByEmail(email string) (*user.UserDomain, *rest_err.RestErr) {
	// Implementação futura
	return nil, nil
}

func (ur *userRepository) FindByID(id string) (*user.UserDomain, *rest_err.RestErr) {
	// Implementação futura
	return nil, nil
}

func (ur *userRepository) Update(id string, userDomain *user.UserDomain) *rest_err.RestErr {
	// Implementação futura
	return nil
}

func (ur *userRepository) Delete(id string) *rest_err.RestErr {
	// Implementação futura
	return nil
}
