package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/auth"
	"github.com/sazbrinna/be_bms/controller/divisi"
	"github.com/sazbrinna/be_bms/controller/location"
	"github.com/sazbrinna/be_bms/controller/profile"
	"github.com/sazbrinna/be_bms/env"
)

func Auth3(r *gin.Engine) {
	router := r.Group("/api/v1/bms/resources")
	router.Use(env.VerifyToken([]string{"admin", "staff"}))

	router.POST("/invitation", auth.Invitation)
	router.POST("/loket/edit", location.EditLocation)
	router.DELETE("/profile/:user_id", profile.DeleteAccount)

	router.PATCH("/divisi/:id_divisi", divisi.UpdateDivisi)
	router.DELETE("/divisi/:id_divisi", divisi.DeleteDivisi)

	router.POST("/region", location.CreateLocation)
	router.DELETE("/region/:id_region", location.DeleteLocation)
}
