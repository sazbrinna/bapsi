package absen

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

func AbsenKeluar(c *gin.Context) {
	var input InputAbsen
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	// absenKeluarDarurat adalah absen yang dilakukan oleh admin

	if len(input.IDUser) > 0 {
		AbsenKeluarDarurat(c, input)
		return
	}

	AbsenKeluarSendiri(c, input)
}

func SaveAkumulasi(profileID string, tambahJam int) {
	// Get current month and year
	now := time.Now()
	year := strconv.Itoa(now.Year())
	monthName := now.Format("January")

	// Check if monthly accumulation data already exists for the profile
	var akumulasi models.AkumulasiBulanan
	if err := models.DB.Where("id_profile = ? AND tahun = ? AND bulan = ? AND is_deleted = ?", profileID, utils.Encrypt(year), utils.Encrypt(monthName), 0).First(&akumulasi).Error; err != nil {
		// Data doesn't exist, create a new record
		akumulasi = models.AkumulasiBulanan{
			IDAkumulasiBulanan: id.NewV6().String(),
			Tahun:              utils.Encrypt(year),
			Bulan:              utils.Encrypt(monthName),
			Akumulasi:          0,
			IsDeleted:          false,
			IDProfile:          profileID,
		}
		models.DB.Create(&akumulasi)
	}

	akumulasi.Akumulasi += tambahJam
	if err := models.DB.Save(&akumulasi).Error; err != nil {
		fmt.Println(err)
	}
}

func JamKeluar() time.Time {

	currentTime := time.Now()

	// Membuat waktu pada pukul 09.00 pada hari yang sama dengan waktu saat ini
	targetTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 16, 0, 0, 0, currentTime.Location())

	// Memeriksa apakah waktu saat ini lebih dari pukul 16.00
	if currentTime.After(targetTime) {
		// Jika iya, set waktu saat ini ke pukul 16.00 pada hari yang sama
		currentTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 16, 0, 0, 0, currentTime.Location())
	}

	return currentTime
}

func AbsenKeluarSendiri(c *gin.Context, input InputAbsen) {
	location, _ := CekLocation(c, input.Latitude, input.Longitude)

	if !location {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "You are out of range"})
		return
	}

	userID, exists := c.Get("id_user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	profileIDInterface, exists := c.Get("id_profile")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "profile not found"})
		return
	}

	profileID, ok := profileIDInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse profile ID"})
		return
	}

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Second)

	var cekAbsensi models.Absensi
	if err := models.DB.Where("id_user = ? AND masuk_at BETWEEN ? AND ? AND is_keluar = ?", userID, startOfDay, endOfDay, 0).First(&cekAbsensi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Anda sudah mengisi presensi keluar",
		})
		return
	}

	cekAbsensi.KeluarAt = JamKeluar()
	cekAbsensi.IsKeluar = true
	cekAbsensi.AkumulasiMenit = int(cekAbsensi.KeluarAt.Sub(cekAbsensi.MasukAt).Minutes())

	models.DB.Save(&cekAbsensi)
	SaveAkumulasi(profileID, cekAbsensi.AkumulasiMenit)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Absen sukses",
	})
}

func AbsenKeluarDarurat(c *gin.Context, input InputAbsen) {
	location, _ := CekLocation(c, input.Latitude, input.Longitude)

	if !location {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "You are out of range"})
		return
	}

	userID := input.IDUser

	var profile models.Profile
	if err := models.DB.
		Select("id_profile").
		Where("id_user = ?", userID).
		First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profileID := profile.IDProfile

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Second)

	var cekAbsensi models.Absensi
	if err := models.DB.Where("id_user = ? AND masuk_at BETWEEN ? AND ? AND is_keluar = ?", userID, startOfDay, endOfDay, 0).First(&cekAbsensi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Anda sudah mengisi presensi keluar",
		})
		return
	}

	cekAbsensi.KeluarAt = JamKeluar()
	cekAbsensi.IsKeluar = true
	cekAbsensi.AkumulasiMenit = int(cekAbsensi.KeluarAt.Sub(cekAbsensi.MasukAt).Minutes())

	models.DB.Save(&cekAbsensi)
	SaveAkumulasi(profileID, cekAbsensi.AkumulasiMenit)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Absen sukses",
	})
}
