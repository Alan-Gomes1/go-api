package controller

import (
	"net/http"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/configuration/validation"
	"github.com/Alan-Gomes1/go-api/src/controller/model/request"
	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userControllerInterface) UpdateUser(c *gin.Context) {
	var userRequest request.UserUpdateRequest
	caller := zap.String("caller", "UpdateUser")
	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, caller)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	_, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		restErr := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)
	restErr := u.service.UpdateUserServices(userId, domain)
	if restErr != nil {
		logger.Error("Error trying to update user", err, caller)
		c.JSON(restErr.Code, err)
		return
	}

	logger.Info(
		"User updated successfully", zap.String("userID", userId), caller,
	)
	c.Status(http.StatusOK)
}
