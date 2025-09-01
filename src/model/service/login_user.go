package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
)

func (u *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (string, *rest_err.Errors) {
	userDomain.EncryptPassword()
	user, err := u.userRepository.FindUserByEmailAndPassword(
		userDomain.GetEmail(), userDomain.GetPassword(),
	)
	if err != nil {
		return "", err
	}
	token, err := user.GenerateToken()
	if err != nil {
		return "", err
	}
	return token, nil
}
