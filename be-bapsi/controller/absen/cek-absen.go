package absen

// ==================================================
// Halaman ini berfungsi untuk
// ==================================================

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/models"
)

func CekAbsenPribadi(c *gin.Context) {
	userID, exists := c.Get("id_user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	var absen models.Absensi
	if err := models.DB.
		Where("id_user = ? AND masuk_at BETWEEN ? AND ?", userID, startOfDay, endOfDay).
		First(&absen).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Not-Check-In"})
	}

	if absen.IsKeluar {
		c.JSON(http.StatusOK, gin.H{"message": "Checked-Out"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checked-In"})
}
