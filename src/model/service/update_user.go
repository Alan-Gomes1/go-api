package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) UpdateUserServices(
	userId string, userDomain model.UserDomainInterface,
) *rest_err.Errors {
	caller := zap.String("caller", "UpdateUserServices")
	err := u.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, caller)
		return err
	}
	return nil
}
