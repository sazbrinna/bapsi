package env

import "github.com/gin-gonic/gin"

func CORSMiddleware(URLEndpoint string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", URLEndpoint)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT, PATCH")

		if URLEndpoint == "local" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
