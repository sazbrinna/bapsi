package env

import (
	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/models"

	"os"
)

type ENV struct {
	PORT string

	// Database
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string

	// Security
	AccessControl string
	ExpectedHost  string
}

func (e ENV) SetEnv() {
	os.Setenv("PORT", e.PORT)
	os.Setenv("DB_HOST", e.DB_HOST)
	os.Setenv("DB_PORT", e.DB_PORT)
	os.Setenv("DB_USERNAME", e.DB_USERNAME)
	os.Setenv("DB_PASSWORD", e.DB_PASSWORD)
	os.Setenv("DB_NAME", e.DB_NAME)
	os.Setenv("AccessControl", e.AccessControl)
	os.Setenv("ExpectedHost", e.ExpectedHost)
}

func (e ENV) RunEnv(r *gin.Engine) {
	models.ConnectToDatabase(e.DB_HOST, e.DB_PORT, e.DB_USERNAME, e.DB_PASSWORD, e.DB_NAME)
	r.Use(CORSMiddleware(e.AccessControl))
	r.Use(SecurityHeader(e.ExpectedHost))
}
