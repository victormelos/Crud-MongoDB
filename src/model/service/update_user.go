package service

import (
	"github.com/victormelos/curso-youtube/src/configuration/logger"
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
)

func (ud *userDomainService) Update(id string, userDomain *user.UserDomain) *rest_err.RestErr {
	logger.Info("Init update user service")

	// Se a senha foi fornecida, criptografa antes de atualizar
	if userDomain.Password != "" {
		if err := userDomain.EncryptPassword(); err != nil {
			logger.Error("Error trying to encrypt password", err)
			return rest_err.NewInternalServerError("Error trying to update user")
		}
	}

	return ud.repository.Update(id, userDomain)
}
