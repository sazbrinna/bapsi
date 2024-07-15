package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/absen"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

func GetAllLocation(c *gin.Context) {
	var locations []models.Region
	if err := models.DB.
		Select("id_region", "region", "latitude", "longitude").
		Where("is_deleted = ?", 0).Find(&locations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var decryptedLocations []gin.H

	for _, loc := range locations {
		lat := absen.ParseFloat(utils.Decrypt(loc.Latitude))
		long := absen.ParseFloat(utils.Decrypt(loc.Longitude))

		latMin := lat - 0.000200
		latMax := lat + 0.000200
		longMin := long - 0.000200
		longMax := long + 0.000200

		decryptedLocation := gin.H{
			"id_region":     loc.IDRegion,
			"region":        utils.Decrypt(loc.Region),
			"latitude":      lat,
			"longitude":     long,
			"latitude_min":  latMin,
			"latitude_max":  latMax,
			"longitude_min": longMin,
			"longitude_max": longMax,
		}
		decryptedLocations = append(decryptedLocations, decryptedLocation)
	}

	c.JSON(http.StatusOK, decryptedLocations)
}
