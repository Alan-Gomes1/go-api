package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
)

func (u *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.Errors) {
	return u.userRepository.FindUserByID(id)
}

func (u *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.Errors) {
	return u.userRepository.FindUserByEmail(email)
}
