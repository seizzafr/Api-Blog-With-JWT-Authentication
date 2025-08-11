package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "api/config"
    "api/models"
	"api/request"
	"api/utils"
	"strconv"
)

// GET /api/categories
func GetCategory(c *gin.Context) {
   pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	var categories []models.Category
	var total int64

	// Hitung total data
	config.DB.Model(&models.Category{}).Count(&total)

	// Ambil data berdasarkan offset & limit
	config.DB.Limit(limit).Offset(offset).Find(&categories)

	// Kirim response paginated
	utils.JSONPaginatedResponse(c, http.StatusOK, categories, total, page, limit, "Data kategori berhasil diambil")
}


// GET /api/categories/:id
func GetCategoryByID(c *gin.Context) {
    var category models.Category
    id := c.Param("id")

    if err := config.DB.First(&category, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
			"status":false,
			"error": "Data not found",
	     	"data":category,
		})
        return
    }

    utils.JSONResponse(c,http.StatusOK,"Data berhasil ditampilkan",category)
}

// POST /api/categories
func CreateCategory(c *gin.Context) {
  var input request.CategoryRequest
if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
        "status":  false,
        "message": "Validasi Input Unvalidated",
        "error":   err.Error(),
        "data":input,
    })
    return
}
// Buat model dari input
category := models.Category{
	Name: input.Name,
}

 config.DB.Create(&category)
 utils.JSONResponse(c,http.StatusOK,"Data berhasil disimpan",category)

}

// PUT /api/categories/:id
func UpdateCategory(c *gin.Context) {
   var input request.CategoryRequest
if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
        "status":  false,
        "message": "Validasi Input Unvalidated",
        "error":   err.Error(),
        "data":input,
    })
    return
}
// Ambil data dari DB dulu
var category models.Category
id := c.Param("id")
  if err := config.DB.First(&category, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
			"status":false,
			"error": "Data not found",
	     	"data":category,
		})
        return
    }

// Update field
category.Name = input.Name
config.DB.Save(&category)

 utils.JSONResponse(c,http.StatusOK,"Data berhasil diupdate",category)
}

// DELETE /api/categories/:id
func DeleteCategory(c *gin.Context) {
    var category models.Category
    id := c.Param("id")

    // 1. Cek apakah kategori ada
    if err := config.DB.First(&category, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "message": "Kategori tidak ditemukan",
            "error":   err.Error(),
            "data":category,
        })
        return
    }

    // 2. Cek apakah masih ada post yang pakai kategori ini
    var posts models.Posts
    if err := config.DB.Where("category_id = ?", category.ID).First(&posts).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Kategori tidak bisa dihapus karena masih digunakan oleh posts",
            "data":posts,
        })
        return
    }

    // 3. Hapus kategori
    if err := config.DB.Delete(&category).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Gagal menghapus kategori",
            "error":   err.Error(),
            "data":category,
        })
        return
    }

     utils.JSONResponse(c,http.StatusOK,"Data berhasil dihapus",category)
}