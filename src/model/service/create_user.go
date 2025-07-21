package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
)

func (u *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.Errors {
	userDomain.EncryptPassword()
	return nil
}
