package service

import (
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomainService) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser service", zap.String("journey", "createUser"))

	if err := ud.UserDomainInterface.EncryptPassword(); err != nil {
		logger.Error("Error trying to encrypt password", err)
		return rest_err.NewInternalServerError("Error trying to create user")
	}

	repo := ud.userRepository
	_, err := repo.CreateUser(ud.UserDomainInterface)
	if err != nil {
		logger.Error("Error trying to call repository", err)
		return err
	}

	return nil
}
