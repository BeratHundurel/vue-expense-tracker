package init_db

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func InitDB() {
	// Define the connection string
	connString := "Server=192.168.1.103,1433;Database=ReStore;integrated security=True;multipleactiveresultsets=True;TrustServerCertificate=True;"

	// Open a connection to the SQL Server
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error connecting to SQL Server:", err.Error())
		return
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error testing SQL Server connection:", err.Error())
		return
	}

	fmt.Println("Connected to SQL Server!")
}
