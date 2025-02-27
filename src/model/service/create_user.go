package service

import (
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) Create(userDomain *user.UserDomain) (*user.UserDomain, *rest_err.RestErr) {
	logger.Info("Init CreateUser service", zap.String("journey", "createUser"))

	if err := userDomain.EncryptPassword(); err != nil {
		logger.Error("Error trying to encrypt password", err)
		return nil, rest_err.NewInternalServerError("Error trying to create user")
	}

	repo := ud.repository
	_, err := repo.Create(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err)
		return nil, err
	}

	return userDomain, nil
}
