package service

import (
	"github.com/victormelos/curso-youtube/src/configuration/rest_err"
	"github.com/victormelos/curso-youtube/src/domain/user"
)

func (ud *userDomainService) FindByEmail(email string) (*user.UserDomain, *rest_err.RestErr) {
	return ud.repository.FindByEmail(email)
}

func (ud *userDomainService) FindByID(id string) (*user.UserDomain, *rest_err.RestErr) {
	return ud.repository.FindByID(id)
}
