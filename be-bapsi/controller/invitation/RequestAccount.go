package invitation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/auth"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

type InputRequestAccount struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	IDDivisi string `json:"id_divisi"`
	IDRegion string `json:"id_region"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func RequestAccount(c *gin.Context) {
	var input InputRequestAccount
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var email models.Users
	if err := models.DB.
		Select("email").
		Where("email = ?", utils.Encrypt(input.Email)).
		First(&email).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email sudah terdaftar",
		})
		return
	}

	var emailRequest models.RequestAccount
	if err := models.DB.
		Select("email").
		Where("email = ?", utils.Encrypt(input.Email)).
		First(&emailRequest).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email sudah melakukan permintaan",
		})
		return
	}

	dbRequestAccount := models.RequestAccount{
		IDRequest: id.NewV6().String(),
		Nama:      utils.Encrypt(input.Nama),
		Email:     utils.Encrypt(input.Email),
		Password:  auth.HashPassword(input.Password),
		IDDivisi:  input.IDDivisi,
		IDRegion:  input.IDRegion,
		Role:      utils.Encrypt(input.Role),
	}

	models.DB.Create(&dbRequestAccount)

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
