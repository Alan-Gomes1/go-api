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
	caller := zap.String("caller", "CreateUser")
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, caller)
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
		logger.Error("Error trying to create user", err, caller)
		c.JSON(err.Code, err)
		return
	}
	userID := zap.String("userID", domainResult.GetID())
	logger.Info("User created successfully", userID, caller)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
