package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/env"
	"github.com/sazbrinna/be_bms/models"
)

func Login(c *gin.Context) {
	var input inputResources
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var account models.Users
	if err := models.DB.Select("id_user", "email", "password", "role").Where("email = ?", utils.Encrypt(input.Email)).First(&account).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
		return
	}

	var profile models.Profile
	if err := models.DB.Select("id_profile").Where("id_user = ?", account.IDUser).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	if !CheckPasswordHash(input.Password, account.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password Salah"})
		return
	}

	tokenJWT, err := env.GenerateToken(account.IDUser, profile.IDProfile, account.Email, account.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokenJWT": tokenJWT,
	})
}
