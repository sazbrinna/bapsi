package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/env"
)

func Auth6(r *gin.Engine) {
	router := r.Group("/api/v1/bms/resources")
	router.Use(env.VerifyToken([]string{"asisten", "staff"}))

}
