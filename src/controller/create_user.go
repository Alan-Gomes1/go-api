package controller

import (
	"fmt"

	"github.com/Alan-Gomes1/go-api/src/configuration/validation"
	"github.com/Alan-Gomes1/go-api/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println("User created successfully:", userRequest)
}
