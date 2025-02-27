package service

import "github.com/victormelos/curso-youtube/src/configuration/rest_err"

func (ud *userDomainService) Delete(id string) *rest_err.RestErr {
	return ud.repository.Delete(id)
}
