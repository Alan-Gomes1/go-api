package repository

import (
	"context"
	"os"

	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
)

const MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"

func (u *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.Errors) {
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := u.databaseConnection.Collection(collectionName)
	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	_id := result.InsertedID.(string)
	userDomain.SetID(_id)
	return userDomain, nil
}
