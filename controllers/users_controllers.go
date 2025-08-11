package controllers

import (
    "net/http"
    "api/config"
    "api/models"
    "api/request"
	"strconv"
    "api/utils"
    "github.com/gin-gonic/gin"
    
)

func GetUsers(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	var users []models.Users
	var total int64

	// Hitung total data
	config.DB.Model(&models.Users{}).Count(&total)
	config.DB.Limit(limit).Offset(offset).Find(&users)
	


	// Kirim response paginated
	utils.JSONPaginatedResponse(c, http.StatusOK, users, total, page, limit, "Data Users berhasil diambil")
}

func GetUsersByID(c *gin.Context) {
    var users models.Users
    id := c.Param("id")

    if err := config.DB.First(&users, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "status":false,
			"error": "User not found",
            "data":users,
        },)
        return
    }

  utils.JSONResponse(c,http.StatusOK,"Data berhasil ditampilkan",users)
}

func CreateUsers(c *gin.Context) {
    var input request.UsersRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
        "status":  false,
        "message": "Validasi Input Unvalidated",
        "error":   err.Error(),
    })
    }

  var existingUser models.Users
    if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
       utils.ValidationErrorResponse(c,"email","Email Sudah Digunakan")
       return
    }

    hashedPassword, err := utils.HashPassword(input.Password)
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
        "status":false,
        "error": "Gagal enkripsi password",
        "data":input.Password,
    })
    return
}

    // Simpan user baru
        users := models.Users{
        Username: input.Username,
        Email:    input.Email,
        Password: hashedPassword,
    }

    config.DB.Create(&users)
    utils.JSONResponse(c,http.StatusOK,"Anda Berhasil Membuat Account",users)
}

func UpdateUsers(c *gin.Context) {
    var users models.Users
    id := c.Param("id")

    // Cek apakah users dengan ID ini ada
    if err := config.DB.First(&users, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  false,
            "message": "User not found",
            "data":users,
        })
        return
    }

    var input request.UsersRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  false,
            "message": "Validasi input unvalidated",
            "error":   err.Error(),
        })
        return
    }

    // Cek apakah email baru sudah digunakan users lain
   var existingUser models.Users
    if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
       utils.ValidationErrorResponse(c,"email","Email Sudah Digunakan")
       return
    }

    // Hash password baru
    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": false,
            "error":  "Gagal enkripsi password",
            "password":input.Password,
        })
        return
    }

    // Update data
    users.Username = input.Username
    users.Email = input.Email
    users.Password = hashedPassword

    if err := config.DB.Save(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": false,
            "error":  "Gagal mengupdate data",
            "data":users,
        })
        return
    }

    utils.JSONResponse(c,http.StatusOK,"Data berhasil diupdate",users)
}


func DeleteUsers(c *gin.Context) {
    var users models.Users
    id := c.Param("id")

    if err := config.DB.First(&users, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{
            "status":false,
			"error": "User not found",
            "data":users,
        },)
        return
    }

    config.DB.Delete(&users)
    utils.JSONResponse(c,http.StatusOK,"Data berhasil dihapus",users)
}



