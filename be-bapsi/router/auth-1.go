package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/divisi"
	"github.com/sazbrinna/be_bms/controller/invitation"
	"github.com/sazbrinna/be_bms/controller/staffs"
	"github.com/sazbrinna/be_bms/env"
)

func Auth1(r *gin.Engine) {
	router := r.Group("/api/v1/bms/resources")
	router.Use(env.VerifyToken([]string{"admin"}))

	router.POST("/admin/invitation", invitation.PostAdminInvitation)

	router.GET("/staffs", staffs.AdmGetAllStaff)

	router.GET("/request", invitation.DaftarRequest)
	router.POST("/request", invitation.IjinkanRequest)
	router.POST("/request/not", invitation.NotIjinkanRequest)

	router.POST("/divisi", divisi.TambahDivisi)
}
