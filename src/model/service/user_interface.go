package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
)

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.Errors
	UpdateUser(string, model.UserDomainInterface) *rest_err.Errors
	FindUser(string) (*model.UserDomainInterface, *rest_err.Errors)
	DeleteUser(string) *rest_err.Errors
}

type userDomainService struct{}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
