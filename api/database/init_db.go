package init_db

import (
	"fmt"
    _ "github.com/microsoft/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	// Define the connection string
	connString := "Server=192.168.1.103,1433;Database=ExpenseTracker;integrated security=True;multipleactiveresultsets=True;TrustServerCertificate=True;"

	// Open a connection to the SQL Server
	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error connecting to SQL Server:", err.Error())
		return
	}
	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error testing SQL Server connection:", err.Error())
		return
	}
	DB = db
	fmt.Println("Connected to SQL Server!")
}
