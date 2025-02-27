package service

import (
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
)

func (ud *userDomainService) Update(id string, userDomain *user.UserDomain) *rest_err.RestErr {
	return ud.repository.Update(id, userDomain)
}
