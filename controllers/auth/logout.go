package auth
import (
    "net/http"
    "strings"
    "api/utils"
    "github.com/gin-gonic/gin"
    
)

func LogoutUser(c *gin.Context) {
	// Ambil token dari header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"error": "Authorization header tidak ditemukan",
			"token":authHeader,
		},)
		return
	}

	// Misal header: "Bearer <token>", kita ambil token-nya
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"error": "Format token tidak valid",
			"token":tokenString,
		},)
		return
	}

	// Masukkan token ke blacklist
	utils.AddToBlacklist(tokenString)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Anda berhasil logout",
		"token":tokenString,
	},)
}
