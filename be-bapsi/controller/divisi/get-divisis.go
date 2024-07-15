package divisi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

func GetDivisi(c *gin.Context) {
	var divisis []models.Divisi
	if err := models.DB.
		Select("id_divisi", "divisi").
		Where("is_deleted = ?", false).
		Find(&divisis).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var decryptedDivisis []gin.H
	for _, divisi := range divisis {
		namaDivisi := utils.Decrypt(divisi.Divisi)

		if namaDivisi == "Admin" || namaDivisi == "Staff" {
			continue
		}

		decryptedDivisi := gin.H{
			"id_divisi": divisi.IDDivisi,
			"divisi":    namaDivisi,
		}

		decryptedDivisis = append(decryptedDivisis, decryptedDivisi)
	}

	c.JSON(http.StatusOK, decryptedDivisis)
}
