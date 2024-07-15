package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/env"
	"github.com/sazbrinna/be_bms/router"
)

func main() {
	port := os.Getenv("PORT")

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// environtment.Environtment()
	env.Environtment(r)
	router.SetupRouter(r)

	fmt.Printf("Server is running on port %s\n", port)
	fmt.Printf("Server is running in: %s\n", pwd)

	r.Run(fmt.Sprintf(":%s", port))
}
