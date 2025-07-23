package controller

import (
	"net/http"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userControllerInterface) DeleteUser(c *gin.Context) {
	caller := zap.String("caller", "DeleteUser")
	userId := c.Param("userId")
	_, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		restErr := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(restErr.Code, restErr)
		return
	}
	restErr := u.service.DeleteUserServices(userId)
	if restErr != nil {
		logger.Error("Error trying to delete user", err, caller)
		c.JSON(restErr.Code, err)
		return
	}

	logger.Info(
		"User deleted successfully", zap.String("userID", userId), caller,
	)
	c.Status(http.StatusOK)
}
