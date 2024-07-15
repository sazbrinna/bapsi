package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

func CreateLocation(c *gin.Context) {
	var input inputEditLocation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	loket := models.Region{
		IDRegion:  id.NewV6().String(),
		Region:    utils.Encrypt(input.Loket),
		Latitude:  utils.Encrypt(input.Latitude),
		Longitude: utils.Encrypt(input.Longitude),
	}

	models.DB.Create(&loket)

	c.JSON(http.StatusCreated, gin.H{"msg": "success"})
}
