package controller

import (
	"net/http"
	"net/mail"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/rest_err"
	"github.com/Alan-Gomes1/go-api/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userControllerInterface) FindUserById(c *gin.Context) {
	userId := c.Param("userId")
	caller := zap.String("caller", "FindUserById")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := "Invalid user id"
		logger.Error(errorMessage, err, caller)
		restErr := rest_err.NewBadRequestError(errorMessage)
		c.JSON(restErr.Code, restErr)
		return
	}

	userDomain, err := u.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to find user", err, caller)
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (u *userControllerInterface) FindUserByEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")
	caller := zap.String("caller", "FindUserByEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := "Invalid user email"
		logger.Error(errorMessage, err, caller)
		restErr := rest_err.NewBadRequestError(errorMessage)
		c.JSON(restErr.Code, err)
		return
	}

	userDomain, err := u.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to find user", err, caller)
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
