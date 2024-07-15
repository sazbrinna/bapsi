package invitation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/models"
)

func NotIjinkanRequest(c *gin.Context) {
	var input InputIjinRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// cari apakah request terdaftar
	var request models.RequestAccount
	if err := models.DB.
		Where("id_request = ?", input.RequestID).
		First(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&request)
}
