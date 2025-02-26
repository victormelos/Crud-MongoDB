package service

import (
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomainService) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser", zap.String("journey", "createUser"))
	logger.Info("Password before encryption", zap.String("password", ud.UserDomainInterface.GetPassword()))

	if err := ud.UserDomainInterface.EncryptPassword(); err != nil {
		logger.Error("Error encrypting password", err)
		return rest_err.NewInternalServerError("Error when trying to encrypt password")
	}

	logger.Info("Password after encryption", zap.String("password", ud.UserDomainInterface.GetPassword()))
	return nil
}
