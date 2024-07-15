package absen

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/models"
)

func KehadiranPesertaBulanan(c *gin.Context) {
	month := c.Param("month")
	year := c.Param("year")

	userID, exists := c.Get("id_user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	firstDay := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.UTC)
	lastDay := time.Date(yearInt, time.Month(monthInt)+1, 0, 0, 0, 0, 0, time.UTC)

	var absens []models.Absensi
	if err := models.DB.
		Select("masuk_at", "keluar_at", "is_keluar").
		Where("id_user = ?", userID).
		Where("is_keluar = ?", true).
		Where("masuk_at BETWEEN ? AND ?", firstDay, lastDay).
		Find(&absens).Error; err != nil {
		if err.Error() != "record not found" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	var absensis []gin.H

	for _, absen := range absens {

		now := absen.MasukAt
		nowKeluar := absen.KeluarAt

		masuk := fmt.Sprintf("%02d.%02d", now.Hour(), now.Minute())
		keluar := fmt.Sprintf("%02d.%02d", nowKeluar.Hour(), nowKeluar.Minute())

		dayName := now.Weekday().String()
		dateFormatted := now.Format("02-01-2006")

		absensi := gin.H{
			"masuk_at":  masuk,
			"keluar_at": keluar,
			"hari":      dayName,
			"tanggal":   dateFormatted,
		}

		absensis = append(absensis, absensi)
	}

	c.JSON(http.StatusOK, absensis)
}

func KehadiranPesertaBulananOther(c *gin.Context) {
	month := c.Param("month")
	year := c.Param("year")

	userID := c.Param("id")

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	firstDay := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.UTC)
	lastDay := time.Date(yearInt, time.Month(monthInt)+1, 0, 0, 0, 0, 0, time.UTC)

	var absens []models.Absensi
	if err := models.DB.
		Select("masuk_at", "keluar_at", "is_keluar").
		Where("id_user = ?", userID).
		Where("is_keluar = ?", true).
		Where("masuk_at BETWEEN ? AND ?", firstDay, lastDay).
		Find(&absens).Error; err != nil {
		fmt.Println(err)
		if err.Error() != "record not found" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	var absensis []gin.H

	for _, absen := range absens {

		now := absen.MasukAt
		nowKeluar := absen.KeluarAt

		masuk := fmt.Sprintf("%02d.%02d", now.Hour(), now.Minute())
		keluar := fmt.Sprintf("%02d.%02d", nowKeluar.Hour(), nowKeluar.Minute())

		dayName := now.Weekday().String()
		dateFormatted := now.Format("02-01-2006")

		absensi := gin.H{
			"masuk_at":  masuk,
			"keluar_at": keluar,
			"hari":      dayName,
			"tanggal":   dateFormatted,
		}

		absensis = append(absensis, absensi)
	}

	c.JSON(http.StatusOK, absensis)
}
