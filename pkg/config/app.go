package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var db *gorm.DB

func Connect() {
	// IMPORTANT: These credentials MUST match your docker run command
	username := "sa"
	password := "MyPassword@123" // Same as SA_PASSWORD in docker run
	host := "localhost"
	port := "1433"

	// Step 1: Connect to master database
	masterConn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=master", username, password, host, port)
	masterDB, err := gorm.Open("mssql", masterConn)
	if err != nil {
		panic(" Cannot connect to SQL Server: " + err.Error())
	}
	defer masterDB.Close()

	// Step 2: Create bookstore database if it doesn't exist
	createDBSQL := "IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'bookstore') CREATE DATABASE bookstore"
	masterDB.Exec(createDBSQL)
	fmt.Println(" Database 'bookstore' ready")

	// Step 3: Connect to bookstore database
	bookstoreConn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=bookstore", username, password, host, port)
	d, err := gorm.Open("mssql", bookstoreConn)
	if err != nil {
		panic(" Cannot connect to bookstore database: " + err.Error())
	}

	db = d
	fmt.Println(" Connected to bookstore database successfully!")
}

func GetDB() *gorm.DB {
	return db
}
