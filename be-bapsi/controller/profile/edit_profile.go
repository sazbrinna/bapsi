package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

type inputProfile struct {
	Nama string `json:"nama"`
}

func EditProfile(c *gin.Context) {
	var input inputProfile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	profileID, exists := c.Get("id_profile")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var profile models.Profile
	if err := models.DB.
		Where("id_profile = ?", profileID).
		First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile.Nama = utils.Encrypt(input.Nama)
	models.DB.Save(&profile)

	c.JSON(http.StatusCreated, gin.H{"msg": "created"})
}
