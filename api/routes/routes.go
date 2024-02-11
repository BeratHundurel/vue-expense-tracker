// routes/routes.go
package routes

import (
	"expense-tracker/go/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/api/helloworld", handlers.Helloworld)
}
