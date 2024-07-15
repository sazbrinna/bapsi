package invitation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

type InputIjinRequest struct {
	RequestID string `json:"request_id"`
}

func IjinkanRequest(c *gin.Context) {
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

	dbUser := models.Users{
		IDUser:   id.NewV6().String(),
		Email:    request.Email,
		Role:     request.Role,
		Password: request.Password,
	}

	dbProfile := models.Profile{
		IDProfile: id.NewV6().String(),
		Nama:      request.Nama,
		IDUser:    dbUser.IDUser,
		IDDivisi:  request.IDDivisi,
		IDRegion:  request.IDRegion,
	}

	models.DB.Create(&dbUser)
	models.DB.Create(&dbProfile)
	models.DB.Delete(&request)
}
