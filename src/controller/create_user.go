package controller

import (
	"net/http"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/validation"
	"github.com/Alan-Gomes1/go-api/src/controller/model/request"
	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/Alan-Gomes1/go-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var UserDomainInterface model.UserDomainInterface

func (u *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validate user info", err,
			zap.String("caller", "CreateUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	domainResult, err := u.service.CreateUserServices(domain)
	if err != nil {
		logger.Error(
			"Error trying to create user", err,
			zap.String("caller", "CreateUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User created successfully",
		zap.String("userID", domainResult.GetID()),
		zap.String("caller", "CreateUser"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
