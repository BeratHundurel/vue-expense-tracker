// routes/routes.go
package routes

import (
	"expense-tracker/go/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/api/transactions", handlers.GetAllTransactions)
	router.POST("/api/addTransaction", handlers.AddTransaction)
	router.DELETE("/api/deleteTransaction/:id", handlers.DeleteTransaction)
}
