package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/absen"
	"github.com/sazbrinna/be_bms/controller/auth"
	"github.com/sazbrinna/be_bms/controller/profile"
	"github.com/sazbrinna/be_bms/env"
)

func Auth7(r *gin.Engine) {
	router := r.Group("/api/v1/bms/resources")
	router.Use(env.VerifyToken([]string{"asisten", "staff", "admin"}))

	router.GET("/absen/cek", absen.CekAbsenPribadi)
	router.GET("/absen/cek/masuk", absen.CekDiRuangan)
	router.POST("/absen/masuk", absen.AbsenMasuk)
	router.POST("/absen/keluar", absen.AbsenKeluar)
	router.GET("profile/absen/:year/:month", absen.KehadiranPesertaBulanan)
	router.GET("profile/absen/:year/:month/:id", absen.KehadiranPesertaBulananOther)

	router.GET("/profile", profile.GetProfile)
	router.POST("/profile", profile.EditProfile)
	router.GET("/email", profile.GetEmail)
	router.GET("/profile/:id", profile.GetProfileOther)

	router.GET("/role", auth.CekRole)

	// lihat Absen
	// get profile
}
