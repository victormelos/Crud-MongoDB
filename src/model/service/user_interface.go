package service

import (
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
)

type userDomainService struct {
	repository user.UserRepositoryInterface
}

func NewUserDomainService(repository user.UserRepositoryInterface) user.UserServiceInterface {
	return &userDomainService{
		repository: repository,
	}
}

func (uds *userDomainService) ValidatePassword(password string) error {
	if len(password) < 6 || len(password) > 100 {
		return rest_err.NewBadRequestError("A senha deve ter entre 6 e 100 caracteres")
	}
	return nil
}
