package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmail(c *gin.Context) {
	Email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": Email})
}
