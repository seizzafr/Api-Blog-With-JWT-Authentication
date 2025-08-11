package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":data,
	})
}

// Response untuk paginated data
type PaginatedData struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}

func JSONPaginatedResponse(c *gin.Context, status int, records interface{}, total int64, page, limit int, message string) {
	totalPages := int((total + int64(limit) - 1) / int64(limit)) // ceil manual

	res := BaseResponse{
		Status:  status,
		Message: message,
		Data: PaginatedData{
			Data:       records,
			Total:      total,
			TotalPages: totalPages,
			Page:       page,
			Limit:      limit,
		},
	}
	c.JSON(status, res)
}

func ValidationErrorResponse(c *gin.Context, message string,field string,) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": false,
		"message": message,
		"errors": gin.H{
			field: message,
		},
	})
}
