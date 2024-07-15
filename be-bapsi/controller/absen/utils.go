package absen

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

type InputAbsen struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Lokasi    string `json:"lokasi"`
	IDUser    string `json:"id_user"`
}

func ParseFloat(input string) float64 {
	value, _ := strconv.ParseFloat(input, 64)
	return value
}

func CekLocation(c *gin.Context, inputLat, inputLong string) (bool, string) {
	var lokasi []models.Region
	if err := models.DB.Find(&lokasi).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return false, ""
	}

	latitude := ParseFloat(inputLat)
	longitude := ParseFloat(inputLong)

	status := false
	var loketAbsen string

	for _, region := range lokasi {
		lat := false
		long := false

		latBapsi := ParseFloat(utils.Decrypt(region.Latitude))
		longBapsi := ParseFloat(utils.Decrypt(region.Longitude))

		latMin := latBapsi - 0.01
		latMax := latBapsi + 0.01
		longMin := longBapsi - 0.01
		longMax := longBapsi + 0.01

		if latitude >= latMin && latitude <= latMax {
			lat = true
		}

		if longitude >= longMin && longitude <= longMax {
			long = true
		}

		if lat && long {
			status = true
			loketAbsen = region.Region
			break
		}
	}

	return status, loketAbsen
}
