package divisi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/models"
)

func DeleteDivisi(c *gin.Context) {
	// using params :id_divisi

	// divisi

	var Divisi models.Divisi
	if err := models.DB.
		Where("id_divisi = ?", c.Param("id_divisi")).
		First(&Divisi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Divisi.IsDeleted = true

	models.DB.Save(&Divisi)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
