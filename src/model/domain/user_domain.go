package domain

import (
	"time"

	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	EncryptPassword() error
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}

type UserRepositoryInterface interface {
	CreateUser(userDomain UserDomainInterface) (UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
