package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
)

func (u *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.Errors) {
	userDomain.EncryptPassword()
	user, err := u.userRepository.FindUserByEmailAndPassword(
		userDomain.GetEmail(), userDomain.GetPassword(),
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
