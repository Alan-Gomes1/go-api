package service

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/Alan-Gomes1/go-api/src/model/repository"
)

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.Errors)
	UpdateUserServices(string, model.UserDomainInterface) *rest_err.Errors
	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.Errors)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.Errors)
	DeleteUserServices(userId string) *rest_err.Errors
}

type userDomainService struct{
	userRepository repository.UserRepository
}

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}
