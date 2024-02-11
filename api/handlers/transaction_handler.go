package handlers

import (
	"expense-tracker/go/models"
	"expense-tracker/go/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllTransactions godoc
// @Summary Get all transactions
// @Description Retrieve a list of all transactions
// @Tags transactions
// @Accept json
// @Produce json
// @Success 200 {array} models.Transaction
// @Router /transactions [get]
func GetAllTransactions(c *gin.Context) {
	transactions, err := services.GetTransactionList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// AddTransaction godoc
// @Summary Add a new transaction
// @Description Add a new transaction to the system
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body models.Transaction true "Transaction object to be added"
// @Success 201 {object} models.Transaction
// @Router /addTransaction [post]
func AddTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.AddTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transaction)
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Delete a transaction from the system
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction Id to be deleted"
// @Success 204
// @Router /deleteTransaction/{id} [delete]
func DeleteTransaction(c *gin.Context) {
	transactionId := c.Param("id")
	id, err := strconv.Atoi(transactionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.DeleteTransaction(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
