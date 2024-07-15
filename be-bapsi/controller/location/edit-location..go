package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

type inputEditLocation struct {
	IDLoket   string `json:"id_loket"`
	Loket     string `json:"loket"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func EditLocation(c *gin.Context) {
	var input inputEditLocation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var loket models.Region
	if err := models.DB.
		Where("id_region = ?", input.IDLoket).
		First(&loket).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	loket.Latitude = utils.Encrypt(input.Latitude)
	loket.Longitude = utils.Encrypt(input.Longitude)
	loket.Region = utils.Encrypt(input.Loket)

	models.DB.Save(&loket)

	c.JSON(http.StatusCreated, gin.H{"msg": "success"})
}
