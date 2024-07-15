package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(host, port, username, password, dbName string) {
	// // DSN without specifying the database name
	// // root:password
	// dsnNoDB := username + ":" + password + "@tcp(" + host + ":" + port + ")/?charset=utf8mb4&parseTime=True&loc=Local"
	// dsnWithDB := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// // Connect to MySQL server
	// database, err := gorm.Open(mysql.Open(dsnNoDB), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect to MySQL server!")
	// }

	// // Create the database if it doesn't exist
	// err = database.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to create database: %v", err))
	// }

	// // Connect to the created database
	// database, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect to the database!")
	// }

	// -------------------------------------------------------------------------
	// Using Sqlite for testing
	// -------------------------------------------------------------------------
	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
