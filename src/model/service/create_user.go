package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.Errors) {
	caller := zap.String("caller", "CreateUserServices")
	user, _ := u.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		logger.Error("Email already registered", nil, caller)
		return nil, rest_err.NewBadRequestError("Email already registered")
	}
	userDomain.EncryptPassword()
	userDomainRepository, err := u.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, caller)
		return nil, err
	}
	return userDomainRepository, nil
}
