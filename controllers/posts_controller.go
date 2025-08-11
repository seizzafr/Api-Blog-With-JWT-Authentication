package controllers

import (
	"os"
    "net/http"
    "github.com/gin-gonic/gin"
    "api/config"
    "api/models"
	"api/request"
	"api/utils"
	"strconv"
    "strings"
	"fmt"
	"path/filepath"
	"time" 
)


// GET /api/posts
func GetPosts(c *gin.Context) {
   pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	var posts []models.Posts
	var total int64

	// Hitung total data
	config.DB.Model(&models.Posts{}).Count(&total)

	// Ambil data berdasarkan offset & limit
	config.DB.Preload("Category").Preload("Users").Limit(limit).Offset(offset).Find(&posts)


	// Kirim response paginated
	utils.JSONPaginatedResponse(c, http.StatusOK, posts, total, page, limit, "Data Posts berhasil diambil")
}


// GET /api/posts/:id
func GetPostsByID(c *gin.Context) {
    var posts models.Posts
    id := c.Param("id")
    config.DB.Preload("Category").Preload("Users").Find(&posts)
     if err := config.DB.First(&posts, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
			"status":false,
			"error": "Data not found",
		})
        return
    }


    utils.JSONResponse(c,http.StatusOK,"Data berhasil ditampilkan",posts)
}

// POST /api/posts
func CreatePosts(c *gin.Context) {
  var input request.PostsRequest
  
if err := c.ShouldBind(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
        "status":  false,
        "message": "Validasi Input Unvalidated",
        "error":   err.Error(),
    })
    return
}


/// Handle upload file
	file, err := c.FormFile("thumbnail")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"error": "File thumbnail wajib diupload"},
		)
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
    allowedExtensions := map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".pdf": true, ".webp": true,
}
if !allowedExtensions[ext] {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":false,
		"error": "Format file harus JPG, PNG, PDF, atau WEBP"},
	)
	return
}

if file.Size > 2<<20 {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":false,
		"error": "Ukuran file maksimal 2MB",		
	})
	return
}


	// Simpan file ke folder uploads/
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	path := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
            "status":false,
			"error": "Gagal simpan file",
		})
		return
	}

	// Buat model
	posts := models.Posts{
		Title:     input.Title,
		Content:   input.Content,
		Thumbnail: "/uploads/" + filename, // simpan path URL-nya
		CategoryID: input.CategoryID,
		UserID:    input.UserID,
	}

if err := config.DB.Create(&posts).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
		"status":false,
		"error": "Gagal simpan ke database",
		"data":posts,
})
    return
}

// ✅ Preload relasi Category dan User agar response lengkap
config.DB.Preload("Category").Preload("Users").First(&posts, posts.ID)

utils.JSONResponse(c, http.StatusOK, "Data berhasil disimpan", posts)



}

// PUT /api/posts/:id
func UpdatePosts(c *gin.Context) {
	var input request.PostsRequest
if err := c.ShouldBind(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
        "status":  false,
        "message": "Validasi Input Unvalidated",
        "error":   err.Error(),
    })
    return
}

	// Ambil data lama dari DB
	var posts models.Posts
	id := c.Param("id")
 if err := config.DB.First(&posts, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
			"status":false,
			"error": "Data not found",
	     	"data":posts,
		})
        return
    }


	// Handle upload file
	file, err := c.FormFile("thumbnail")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"error": "File thumbnail wajib diupload"},
		)
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExtensions := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".pdf": true, ".webp": true,
	}
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"error": "Format file harus JPG, PNG, PDF, atau WEBP",
			"data":gin.H{
				"Extension File":ext,
			},
	})
		return
	}

	if file.Size > 2<<20 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"error": "Ukuran file maksimal 2MB"})
		return
	}

	// Simpan file ke folder uploads/
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	path := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":false,
			"error": "Gagal simpan file"})
		return
	}

	// Update data posts
	posts.Title = input.Title
	posts.Content = input.Content
	posts.Thumbnail = "/uploads/" + filename
	posts.CategoryID = input.CategoryID
	posts.UserID = input.UserID

	// Simpan ke DB
	if err := config.DB.Save(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":false,
			"error": "Gagal update data",
			"data":posts,
		},
		)
		return
	}

	// Preload Category dan User agar lengkap saat response
	config.DB.Preload("Category").Preload("Users").First(&posts, posts.ID)
    utils.JSONResponse(c, http.StatusOK, "Data berhasil diupdate", posts)
}





func DeletePosts(c *gin.Context) {
    var posts models.Posts
    id := c.Param("id")

    // Cari data post berdasarkan ID
    if err := config.DB.Preload("Category").Preload("Users").First(&posts, id).Error; err != nil {
        utils.JSONResponse(c, http.StatusNotFound, "Data tidak ditemukan", nil)
        return
    }

    // Hapus file thumbnail jika ada
    if posts.Thumbnail != "" {
        path := posts.Thumbnail
        if err := os.Remove(path); err != nil {
            fmt.Println("Gagal hapus file:", err)
        }
    }

    // Hapus data post dari database
    config.DB.Delete(&posts)

    utils.JSONResponse(c, http.StatusOK, "Data berhasil dihapus", posts)
}
