package absen

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

func AbsenMasuk(c *gin.Context) {
	var input InputAbsen
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	if len(input.IDUser) > 0 {
		AbsenMasukDarurat(c, input)
		return
	}

	AbsenMasukSendiri(c, input)
}

func JamMasuk() time.Time {
	currentTime := time.Now()

	// Membuat waktu pada pukul 09.00 pada hari yang sama dengan waktu saat ini
	targetTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 9, 0, 0, 0, currentTime.Location())

	// Memeriksa apakah waktu saat ini kurang dari pukul 09.00
	if currentTime.Before(targetTime) {
		// Jika iya, set waktu ke pukul 09.00
		currentTime = targetTime
	}

	return currentTime
}

func AbsenMasukDarurat(c *gin.Context, input InputAbsen) {
	location, loket := CekLocation(c, input.Latitude, input.Longitude)

	if !location {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "You are out of range"})
		return
	}

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Second)

	userID := input.IDUser

	var user models.Users
	if err := models.DB.
		Select("id_user").
		Where("id_user = ?", userID).
		First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User tidak terdaftar"})
		return
	}

	var statusAbsen models.Absensi
	if err := models.DB.Where("id_user = ? AND masuk_at BETWEEN ? AND ? AND is_keluar = ?", userID, startOfDay, endOfDay, 0).First(&statusAbsen).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "User sudah presensi masuk hari ini",
		})
		return
	}

	dbAbsen := models.Absensi{
		IDAbsensi: id.NewV6().String(),
		IDUser:    userID,
		MasukAt:   JamMasuk(),
		Lokasi:    loket,
	}

	if err := models.DB.Create(&dbAbsen).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Absen sukses",
	})
}

func AbsenMasukSendiri(c *gin.Context, input InputAbsen) {
	location, loket := CekLocation(c, input.Latitude, input.Longitude)

	if !location {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "You are out of range"})
		return
	}

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Second)

	userIDInterface, exists := c.Get("id_user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	var user models.Users
	if err := models.DB.
		Select("id_user").
		Where("id_user = ?", userID).
		First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User tidak terdaftar"})
		return
	}

	var statusAbsen models.Absensi
	if err := models.DB.
		Where("id_user = ? AND masuk_at BETWEEN ? AND ? AND is_keluar = ?", userID, startOfDay, endOfDay, 0).
		First(&statusAbsen).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "User sudah presensi masuk hari ini",
		})
		return
	}

	dbAbsen := models.Absensi{
		IDAbsensi: id.NewV6().String(),
		IDUser:    userID,
		MasukAt:   JamMasuk(),
		Lokasi:    loket,
	}

	if err := models.DB.Create(&dbAbsen).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Absen sukses",
	})
}
