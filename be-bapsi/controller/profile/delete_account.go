package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/models"
)

func DeleteAccount(c *gin.Context) {
	userID := c.Param("user_id")

	var user models.Users
	if err := models.DB.
		Where("id_user = ?", userID).
		First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var profile models.Profile
	if err := models.DB.
		Where("id_user = ?", userID).
		First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	profile.IsDeleted = true
	user.IsDeleted = true

	models.DB.Save(&profile)
	models.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
