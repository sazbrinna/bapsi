package env

import (
	"github.com/gin-gonic/gin"
)

func Environtment(r *gin.Engine) {
	myEnv := ENV{
		PORT: "8000",

		DB_HOST:     "172.42.0.4",
		DB_PORT:     "3306",
		DB_USERNAME: "root",
		DB_PASSWORD: "faizal",
		DB_NAME:     "bapsi_management_system",

		AccessControl: "*",
		ExpectedHost:  "https://bapsi-back.engkol.cloud",
	}

	myEnv.RunEnv(r)
	myEnv.SetEnv()
}
