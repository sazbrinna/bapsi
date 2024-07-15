package divisi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

func UpdateDivisi(c *gin.Context) {
	// using params :id_divisi

	// divisi

	var input InputDivisi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(input)

	var Divisi models.Divisi
	if err := models.DB.
		Where("id_divisi = ?", c.Param("id_divisi")).
		First(&Divisi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Divisi.Divisi = utils.Encrypt(input.Divisi)

	models.DB.Save(&Divisi)

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}
