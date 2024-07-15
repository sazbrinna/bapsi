package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	Auth0(r)
	Auth1(r)
	Auth2(r)
	Auth3(r)
	Auth4(r)
	Auth5(r)
	Auth6(r)
	Auth7(r)
}
