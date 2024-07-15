package main

import (
	"github.com/sazbrinna/be_bms/models"
	"github.com/sazbrinna/be_bms/models/migration"
)

type ENV struct {
	PORT string

	// Database
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

func (e ENV) RunDatabase() {
	models.ConnectToDatabase(e.DB_HOST, e.DB_PORT, e.DB_USERNAME, e.DB_PASSWORD, e.DB_NAME)
}

func main() {
	myEnv := ENV{
		PORT: "8000",

		DB_HOST:     "127.20.0.2",
		DB_PORT:     "3306",
		DB_USERNAME: "root",
		DB_PASSWORD: "faizal",
		DB_NAME:     "bapsi_management_system",
	}

	myEnv.RunDatabase()
	migration.TestDecrypt()
}
