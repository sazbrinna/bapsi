package invitation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

type InputInvitation struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	Division string `json:"division"`
	Region   string `json:"region"`
}

func PostAdminInvitation(c *gin.Context) {
	var input InputInvitation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var division string

	switch input.Role {
	case "asisten":
		division = "asisten"
	case "admin":
		division = "admin"
	default:
		division = input.Role
	}

	// cek email availability
	// becarefull if you wana edit the error status message because it use on frontend to make decision

	var emailUser models.Users
	if err := models.DB.
		Select("email").
		Where("email = ?", utils.Encrypt(input.Email)).
		First(&emailUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email invalid"})
		return
	}

	var emailInvitation models.Invitation
	if err := models.DB.
		Select("email").
		Where("email = ?", utils.Encrypt(input.Email)).
		First(&emailInvitation).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already invited"})
		return
	}

	// save data to the database
	dataInvitation := models.Invitation{
		IDInvitation: id.NewV6().String(),
		Email:        utils.Encrypt(input.Email),
		Divisi:       utils.Encrypt(division),
		Role:         utils.Encrypt(input.Role),
		Region:       utils.Encrypt(input.Region),
	}

	models.DB.Create(&dataInvitation)

	c.JSON(http.StatusCreated, gin.H{"message": "success created"})
}
