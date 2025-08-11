package auth

import (
    "net/http"
    "api/config"
    "api/models"
    "api/utils"
    "github.com/gin-gonic/gin"
    
)
func LoginUser(c *gin.Context) {
	var input models.Users
	var user models.Users

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
        "status":  false,
        "message": "Validasi Input Unvalidated",
        "error":   err.Error(),
    })
		return
	}

	// Cek apakah user ada
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":false,
			"error": "Email tidak ditemukan",
			"email":input.Email,
		},)
		return
	}

	// Cek password
	if err := utils.CheckPassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":false,
			"error": "Password salah",
			"password":input.Password,
		},)
		return
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":false,
			"error": "Gagal membuat token",
			"token":token,
		},)
		return
	}

	c.JSON(http.StatusOK, gin.H{
    "status":true,
    "message":"Anda Berhasil Login",
	"user": gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	},
     "token": token,
})

}
