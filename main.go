package main

import (
	"context"
	"log"

	"github.com/Alan-Gomes1/go-api/src/configuration/database/mongodb"
	"github.com/Alan-Gomes1/go-api/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	userController := initDependencies(database)
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
