package repository

import (
	"context"
	"os"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/Alan-Gomes1/go-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userRepository) UpdateUser(
	userId string, userDomain model.UserDomainInterface,
) *rest_err.Errors {
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := u.databaseConnection.Collection(collectionName)
	objectId, _ := primitive.ObjectIDFromHex(userId)
	value := converter.ConvertDomainToEntity(userDomain)
	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: value}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		caller := zap.String("caller", "userRepository.UpdateUser")
		logger.Error("Error trying to update user", err, caller)
		return rest_err.NewInternalServerError(err.Error())
	}
	return nil
}
