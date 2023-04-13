package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database is a global variable representing the connection to the database.
var Database *gorm.DB

// urlDSN is a global variable representing the connection string for the database.
var urlDSN = "root:Jay@7000@tcp(localhost:3306)/mydb"

// err is a global variable representing any error occurred during the database connection.
var err error

// DataMigration is a function that creates necessary database tables and columns for the Employee model.
func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed :(")
	}

	err = Database.AutoMigrate(&Employee{})
	if err != nil {
		fmt.Print(err.Error())
		panic("migration failed :(")
	}
}
