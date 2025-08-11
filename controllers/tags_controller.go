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

// GET /api/Tags
func GetTags(c *gin.Context) {
   pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	var Tags []models.Tags
	var total int64

	// Hitung total data
	config.DB.Model(&models.Tags{}).Count(&total)

	// Ambil data berdasarkan offset & limit
	config.DB.Limit(limit).Offset(offset).Find(&Tags)

	// Kirim response paginated
	utils.JSONPaginatedResponse(c, http.StatusOK, Tags, total, page, limit, "Data tag berhasil diambil")
}


// GET /api/Tags/:id
func GetTagsByID(c *gin.Context) {
    var tags models.Tags
    id := c.Param("id")

    if err := config.DB.First(&tags, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "status":false,
            "error": "Data not found",
            "data":tags,
        },
        )
        return
    }

    utils.JSONResponse(c,http.StatusOK,"Data berhasil ditampilkan",tags)
}

// POST /api/Tags
func CreateTags(c *gin.Context) {
  var input request.TagsRequest
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
    tags := models.Tags{
	Name: input.Name,
}

 config.DB.Create(&tags)
 utils.JSONResponse(c,http.StatusOK,"Data berhasil disimpan",tags)

}

// PUT /api/Tags/:id
func UpdateTags(c *gin.Context) {
   var input request.TagsRequest
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
var tags models.Tags
id := c.Param("id")
if err := config.DB.First(&tags, id).Error; err != nil {
	        c.JSON(http.StatusNotFound, gin.H{
            "status":false,
            "error": "Data not found",
            "data":tags,
        },)
	return
}

// Update field
tags.Name = input.Name
config.DB.Save(&tags)

 utils.JSONResponse(c,http.StatusOK,"Data berhasil diupdate",tags)
}

// DELETE /api/Tags/:id
func DeleteTags(c *gin.Context) {
    var tags models.Tags
    id := c.Param("id")

    if err := config.DB.First(&tags, id).Error; err != nil {
              c.JSON(http.StatusNotFound, gin.H{
            "status":false,
            "error": "Data not found",
            "data":tags,
        },)
        return
    }

    config.DB.Delete(&tags)
    utils.JSONResponse(c,http.StatusOK,"Data berhasil dihapus",tags)
}
