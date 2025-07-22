package repository

import (
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.Errors)
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{databaseConnection}
}
