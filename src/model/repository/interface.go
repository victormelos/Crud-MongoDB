package repository

import (
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/model/service"
)

type UserRepository interface {
	CreateUser(userDomain service.UserDomainInterface) (service.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (service.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (service.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, service.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
