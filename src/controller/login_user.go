package controller

import (
	"net/http"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"github.com/Alan-Gomes1/go-api/src/configuration/validation"
	"github.com/Alan-Gomes1/go-api/src/controller/model/request"
	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *userControllerInterface) LoginUser(c *gin.Context) {
	caller := zap.String("caller", "LoginUser")
	var userRequest request.UserLoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, caller)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
	token, err := u.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call LoginUserServices", err, caller)
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": "Bearer " + token})
}
