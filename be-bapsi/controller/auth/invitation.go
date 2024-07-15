package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

func Invitation(c *gin.Context) {
	var input inputResources
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dbInvitation := models.Invitation{
		IDInvitation: id.NewV6().String(),
		Email:        utils.Encrypt(input.Email),
		Role:         utils.Encrypt(input.Role),
	}

	models.DB.Create(&dbInvitation)

	c.JSON(http.StatusCreated, gin.H{
		"email": utils.Decrypt(dbInvitation.Email),
		"role":  utils.Decrypt(dbInvitation.Role),
	})
}
