package env

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

// from another project
// im not sure this is correct or not

var secretKey = []byte("G%3jyF83%?9Bg7,uX;(g-}tug:0n!IFPeJ{7qK!s>@_MGmbl!t/Y7:/j+T|hIZR")

type MyClaims struct {
	jwt.StandardClaims
	IDUser    string `json:"id_user"`
	IDProfile string `json:"id_profile"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

func GenerateToken(id, profile_id, email, role string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 3).Unix(), // Token expires in 1 minute
		},
		IDUser:    id,
		IDProfile: profile_id,
		Email:     email,
		Role:      role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(access []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString = parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		// Extract claims from token
		claims, ok := token.Claims.(*MyClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		var user models.Users
		if err := models.DB.Select("id_user").Where("id_user = ?", claims.IDUser).Find(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Check if the user's role is allowed to access the resource
		allowed := false
		for _, role := range access {
			if utils.Encrypt(role) == claims.Role {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// User has access, proceed with the request
		c.Set("email", claims.Email)
		c.Set("id_user", claims.IDUser)
		c.Set("id_profile", claims.IDProfile)
		c.Set("role", claims.Role)
		c.Next()
	}
}
