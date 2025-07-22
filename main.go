package main

import (
	"log"

	"github.com/Alan-Gomes1/go-api/src/configuration/database/mongodb"
	"github.com/Alan-Gomes1/go-api/src/controller"
	"github.com/Alan-Gomes1/go-api/src/controller/routes"
	"github.com/Alan-Gomes1/go-api/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.NewMongoDBConnection()
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
