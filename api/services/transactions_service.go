package services

import (
	database "expense-tracker/go/database"
	"expense-tracker/go/models"
)

func GetTransactionList() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := database.DB.Select(&transactions, "SELECT * FROM Transactions")
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func AddTransaction(transaction models.Transaction) error {
	_, err := database.DB.NamedExec("INSERT INTO Transactions (Text, Amount) VALUES (:Text, :Amount)", transaction)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTransaction(id int) error {
	_, err := database.DB.NamedExec("DELETE FROM Transactions WHERE Id=:id", map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	return nil
}
