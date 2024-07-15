package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

// This is for confirmation in registration

func GetInvitation(c *gin.Context) {
	email := c.Param("email")

	var invitation models.Invitation
	if err := models.DB.
		Where("email = ?", utils.Encrypt(email)).
		First(&invitation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"eror": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"divisi": utils.Decrypt(invitation.Divisi),
		"region": utils.Decrypt(invitation.Region),
		"role":   utils.Decrypt(invitation.Role),
	})
}
