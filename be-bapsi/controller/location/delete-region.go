package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/models"
)

func DeleteLocation(c *gin.Context) {
	// params: /location/delete/:id_region

	var loket models.Region
	if err := models.DB.
		Where("id_region = ?", c.Param("id_region")).
		First(&loket).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	loket.IsDeleted = true

	models.DB.Save(&loket)

	c.JSON(http.StatusCreated, gin.H{"msg": "success"})
}
