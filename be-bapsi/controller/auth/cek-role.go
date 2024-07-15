package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
)

func CekRole(c *gin.Context) {
	roleInterface, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	role, ok := roleInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse profile ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": utils.Decrypt(role)})
}
