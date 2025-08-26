package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *userDomainService) DeleteUserServices(userId string) *rest_err.Errors {
	caller := zap.String("caller", "DeleteUserServices")
	err := u.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository", err, caller)
		return err
	}
	return nil
}
