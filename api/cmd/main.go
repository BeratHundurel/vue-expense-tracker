// main.go
package main

import (
	docs "expense-tracker/go/cmd/docs"
	database "expense-tracker/go/database"
	"expense-tracker/go/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	port = ":8080"
)

func main() {
	// Initialize the database.
	database.InitDB()
	// Set up Swagger documentation
	docs.SwaggerInfo.BasePath = "/api"
	// Set up Gin router
	router := gin.Default()
	// Set up routes
	routes.SetRoutes(router)
	// Serve Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Run the server
	router.Run(port)
}
