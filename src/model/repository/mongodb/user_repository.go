package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseConnection *mongo.Client
}

func NewUserRepository(database *mongo.Client) domain.UserRepositoryInterface {
	return &userRepository{
		databaseConnection: database,
	}
}

func (ur *userRepository) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.databaseConnection.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	userDomain.SetCreatedAt(time.Now())
	userDomain.SetUpdatedAt(time.Now())

	logger.Info("Attempting to insert user into MongoDB")
	_, err := collection.InsertOne(context.Background(), userDomain)
	if err != nil {
		logger.Error("Error inserting user into MongoDB", err)
		return nil, rest_err.NewInternalServerError("Error trying to create user")
	}

	logger.Info("User successfully inserted into MongoDB")
	return userDomain, nil
}

func (ur *userRepository) FindUserByEmail(email string) (domain.UserDomainInterface, *rest_err.RestErr) {
	// Implementação futura
	return nil, nil
}

func (ur *userRepository) FindUserByID(id string) (domain.UserDomainInterface, *rest_err.RestErr) {
	// Implementação futura
	return nil, nil
}

func (ur *userRepository) UpdateUser(id string, userDomain domain.UserDomainInterface) *rest_err.RestErr {
	// Implementação futura
	return nil
}

func (ur *userRepository) DeleteUser(id string) *rest_err.RestErr {
	// Implementação futura
	return nil
}
