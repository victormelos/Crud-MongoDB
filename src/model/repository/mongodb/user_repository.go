package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/model/repository"
	"github.com/victormelos/curso-youtube/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseConnection *mongo.Client
}

func NewUserRepository(database *mongo.Client) repository.UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

func (ur *userRepository) CreateUser(userDomain service.UserDomainInterface) (service.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.databaseConnection.Database(os.Getenv("MONGODB_USER_DB")).Collection("users")

	userDomain.SetCreatedAt(time.Now())
	userDomain.SetUpdatedAt(time.Now())

	_, err := collection.InsertOne(context.Background(), userDomain)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to create user")
	}

	return userDomain, nil
}

func (ur *userRepository) FindUserByEmail(email string) (service.UserDomainInterface, *rest_err.RestErr) {
	// Implementação futura
	return nil, nil
}

func (ur *userRepository) FindUserByID(id string) (service.UserDomainInterface, *rest_err.RestErr) {
	// Implementação futura
	return nil, nil
}

func (ur *userRepository) UpdateUser(id string, userDomain service.UserDomainInterface) *rest_err.RestErr {
	// Implementação futura
	return nil
}

func (ur *userRepository) DeleteUser(id string) *rest_err.RestErr {
	// Implementação futura
	return nil
}
