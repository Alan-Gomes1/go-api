package routes

import (
	"github.com/Alan-Gomes1/go-api/src/controller"
	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/user/:userId", model.VerifyTokenMIddleware, userController.FindUserById)
	r.GET("/userByEmail/:userEmail", model.VerifyTokenMIddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
}
