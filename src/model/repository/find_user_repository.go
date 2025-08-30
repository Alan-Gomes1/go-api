package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/Alan-Gomes1/go-api/src/model/repository/entity"
	"github.com/Alan-Gomes1/go-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *userRepository) FindUserByID(id string) (
	model.UserDomainInterface, *rest_err.Errors,
) {
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := u.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with id %s not found", id)
			logger.Error(
				errorMessage, err,
				zap.String("caller", "userRepository.FindUserByID"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(
			errorMessage, err,
			zap.String("caller", "userRepository.FindUserByID"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info(
		"User found successfully",
		zap.String("userID", userEntity.ID.Hex()),
		zap.String("caller", "userRepository.FindUserByID"),
	)
	userDomain := converter.ConvertEntityToDomain(*userEntity)
	return userDomain, nil
}

func (u *userRepository) FindUserByEmail(email string) (
	model.UserDomainInterface, *rest_err.Errors,
) {
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := u.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with email %s not found", email)
			logger.Error(
				errorMessage, err,
				zap.String("caller", "userRepository.FindUserByEmail"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(
			errorMessage, err,
			zap.String("caller", "userRepository.FindUserByEmail"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (u *userRepository) FindUserByEmailAndPassword(email, password string) (
	model.UserDomainInterface, *rest_err.Errors,
) {
	caller := zap.String("caller", "userRepository.FindUserByEmailAndPassword")
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := u.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password invalid"
			logger.Error(errorMessage, err, caller)
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage, err, caller)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil
}
