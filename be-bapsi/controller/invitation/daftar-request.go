package invitation

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

func DaftarRequest(c *gin.Context) {
	var request []models.RequestAccount
	if err := models.DB.
		Find(&request).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var decryptedValues []gin.H
	for _, value := range request {

		var region models.Region
		if err := models.DB.
			Select("region").
			Where("id_region = ?", value.IDRegion).
			First(&region).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var divisi models.Divisi
		if err := models.DB.
			Select("divisi").
			Where("id_divisi = ?", value.IDDivisi).
			First(&divisi).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("and here")

		decryptedValue := gin.H{
			"id_request": value.IDRequest,
			"region":     utils.Decrypt(region.Region),
			"divisi":     utils.Decrypt(divisi.Divisi),
			"role":       utils.Decrypt(value.Role),
			"nama":       utils.Decrypt(value.Nama),
			"email":      utils.Decrypt(value.Email),
		}

		decryptedValues = append(decryptedValues, decryptedValue)
	}

	c.JSON(http.StatusOK, decryptedValues)
}
