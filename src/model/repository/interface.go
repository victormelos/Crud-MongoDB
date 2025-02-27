package repository

import (
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
)

type UserRepository interface {
	Create(userDomain *user.UserDomain) (*user.UserDomain, *rest_err.RestErr)
	FindByEmail(email string) (*user.UserDomain, *rest_err.RestErr)
	FindByID(id string) (*user.UserDomain, *rest_err.RestErr)
	Update(id string, user *user.UserDomain) *rest_err.RestErr
	Delete(id string) *rest_err.RestErr
}
