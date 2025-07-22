package controller

import (
	"github.com/Alan-Gomes1/go-api/src/model/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

func NewUserControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{service}
}
