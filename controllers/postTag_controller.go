package controllers

import (
    "api/config" // ganti sesuai nama module kamu
    "api/models"
    "api/utils"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// GET /api/post-tags
func GetPostTags(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")

    page, _ := strconv.Atoi(pageStr)
    limit, _ := strconv.Atoi(limitStr)
    offset := (page - 1) * limit

    var postTags []models.PostTag
    var total int64

    config.DB.Model(&models.PostTag{}).Count(&total)
    config.DB.
    Preload("Posts.Users").
    Preload("Posts.Category").
    Preload("Tags").
    Limit(limit).
    Offset(offset).
    Find(&postTags)


    utils.JSONPaginatedResponse(c, http.StatusOK, postTags, total, page, limit, "Data PostTag berhasil diambil")
}

// GET /api/post-tags/:id
func GetPostTagByID(c *gin.Context) {
    postID := c.Param("post_id")
    tagID := c.Param("tag_id")

    var postTag models.PostTag

    err := config.DB.
        Preload("Posts.Users").
        Preload("Posts.Category").
        Preload("Tags").
        Where("post_id = ? AND tag_id = ?", postID, tagID).
        First(&postTag).Error

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "status":false,
            "error": "Data tidak ditemukan",
            "data":postTag,
        },)
        return
    }


    utils.JSONResponse(c, http.StatusOK, "Data berhasil ditampilkan", postTag)
}

// POST /api/post-tags
func CreatePostTag(c *gin.Context) {
    var input models.PostTag
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validasi: pastikan Post dan Tag-nya ada
    var posts models.Posts
    if err := config.DB.First(&posts, input.PostID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "status":false,
            "error": "Posts tidak ditemukan",
            "data":posts,
        },
        )
        return
    }

    var tags models.Tags
    if err := config.DB.First(&tags, input.TagID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H {
            "status":false,
            "error": "Tag tidak ditemukan",
            "data":tags,
        },
        )
        return
    }

    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

 var result models.PostTag
    if err := config.DB.
        Where("post_id = ? AND tag_id = ?", input.PostID, input.TagID).
        Preload("Posts.Users").
        Preload("Posts.Category"). 
        Preload("Tags").
        First(&result).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    utils.JSONResponse(c, http.StatusOK, "Data berhasil disimpan", result)
}

// DELETE /api/post-tags
// func DeletePostTag(c *gin.Context) {
//     postID := c.Param("post_id")
//     tagID := c.Param("tag_id")

//     fmt.Println("Post ID:", postID)
// 	fmt.Println("Tag ID:", tagID)
//     var postTag models.PostTag
    


//     utils.JSONResponse(c, http.StatusOK, "Data berhasil dihapus")
// }
