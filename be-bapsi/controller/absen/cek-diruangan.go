package absen

// fungsinya adalah untuk mengambil data siapa saja yang sedang dalam ruangan pada hari itu

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

func CekDiRuangan(c *gin.Context) {
	var absens []models.Absensi

	// Get start and end of today
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Second) // end of today

	err := models.DB.
		Select("id_user", "lokasi", "masuk_at", "keluar_at", "is_keluar").
		Where("masuk_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Find(&absens).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	var emps []gin.H
	for _, absen := range absens {
		var user models.Profile
		err := models.DB.
			Select("id_user", "nama").
			Where("id_user = ?", absen.IDUser).
			First(&user).Error
		if err != nil {
			continue
		}

		now := absen.MasukAt
		pulang := absen.KeluarAt
		formattedTime := now.Format("15:04")
		formattedPulang := pulang.Format("15:04")

		emp := gin.H{
			"id_user":      user.IDUser,
			"nama":         utils.Decrypt(user.Nama),
			"lokasi_absen": utils.Decrypt(absen.Lokasi),
			"jam_absen":    formattedTime,
			"jam_pulang":   formattedPulang,
			"is_keluar":    absen.IsKeluar,
		}
		emps = append(emps, emp)
	}

	c.JSON(http.StatusOK, emps)
}
