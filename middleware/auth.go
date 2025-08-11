package middleware

import (
	"net/http"
	"os"
	"strings"
	"api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":false,
				"error": "Token tidak ditemukan",
				"token":
					authHeader,
			},)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":false,
				"error": "Token tidak valid",
				"token":token,
			},)
			c.Abort()
			return
		}

		if utils.IsTokenBlacklisted(tokenString) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":false,
			"error": "Token sudah tidak berlaku (logout)",
			"token":tokenString,
		},)
		return
}


		claims := token.Claims.(jwt.MapClaims)

		// Simpan informasi user di context
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("email", claims["email"].(string))

		c.Next()

	}
}
