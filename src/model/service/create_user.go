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
	userDomain.EncryptPassword()
	userDomainRepository, err := u.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error(
			"Error trying to call repository", err,
			zap.String("caller", "CreateUserServices"),
		)
		return nil, err
	}
	return userDomainRepository, nil
}
