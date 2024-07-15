package router

// ini adalah halaman untuk router
// nilai router 1 = admin
// nilai router 2 = untuk staff
// nilai router 4 = untuk asisten

import (
	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/auth"
	"github.com/sazbrinna/be_bms/controller/divisi"
	"github.com/sazbrinna/be_bms/controller/invitation"
	"github.com/sazbrinna/be_bms/controller/location"
)

func Auth0(r *gin.Engine) {
	router := r.Group("/api/v1/bms/resources")

	router.POST("/login", auth.Login)
	// router.POST("/registration", auth.Registration)
	router.POST("/registration", invitation.RequestAccount)
	router.GET("/invitation/:email", auth.GetInvitation)

	router.GET("/location/all", location.GetAllLocation)

	router.GET("/divisi", divisi.GetDivisi)

}
