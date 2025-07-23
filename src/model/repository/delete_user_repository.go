package repository

import (
	"context"
	"os"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userRepository) DeleteUser(userId string) *rest_err.Errors {
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := u.databaseConnection.Collection(collectionName)
	objectId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objectId}}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		caller := zap.String("caller", "userRepository.DeleteUser")
		logger.Error("Error trying to delete user", err, caller)
		return rest_err.NewInternalServerError(err.Error())
	}
	return nil
}
