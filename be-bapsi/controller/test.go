package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello",
	})
}

func TestAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Sukses",
	})
}
