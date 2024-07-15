package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/env"
	"github.com/sazbrinna/be_bms/models"
)

func Registration(c *gin.Context) {
	var input inputResources
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// ----------------------------------------------------------------
	// Check Invitation berfore registration
	// ----------------------------------------------------------------

	var invitation models.Invitation
	if err := models.DB.Where("email = ?", utils.Encrypt(input.Email)).First(&invitation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var region models.Region
	err := models.DB.
		Select("id_region").
		Where("region = ?", invitation.Region).
		First(&region).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var divisi models.Divisi
	if err := models.DB.
		Select("id_divisi").
		Where("divisi = ?", invitation.Divisi).
		First(&divisi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	// ----------------------------------------------------------------
	// Save to the database
	// ----------------------------------------------------------------

	dbUser := models.Users{
		IDUser:   id.NewV6().String(),
		Email:    utils.Encrypt(input.Email),
		Password: HashPassword(input.Password),
		Role:     invitation.Role,
	}

	dbProfile := models.Profile{
		IDProfile: id.NewV6().String(),
		IDUser:    dbUser.IDUser,
		IDDivisi:  divisi.IDDivisi,
		IDRegion:  region.IDRegion,
		Nama:      utils.Encrypt(input.Nama),
	}

	if err := models.DB.Create(&dbUser).Error; err != nil {
		fmt.Println(err.Error())
	}

	models.DB.Create(&dbProfile)
	models.DB.Delete(&invitation)

	// generate token jwt
	tokenJWT, err := env.GenerateToken(dbUser.IDUser, dbProfile.IDProfile, dbUser.Email, dbUser.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokenJWT": tokenJWT,
		"email":    utils.Decrypt(dbUser.Email),
	})
}
