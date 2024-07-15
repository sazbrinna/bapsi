package divisi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

type InputDivisi struct {
	Divisi string `json:"divisi"`
	Group  string `json:"group"`
}

func TambahDivisi(c *gin.Context) {
	var input InputDivisi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dbDivisi := models.Divisi{
		IDDivisi: id.NewV6().String(),
		Divisi:   utils.Encrypt(input.Divisi),
		Group:    utils.Encrypt(input.Group),
	}

	models.DB.Create(&dbDivisi)

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}
