package response

import "github.com/gin-gonic/gin"

type Pagination struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

func SuccessList(c *gin.Context, data interface{}, pagination Pagination) {
	c.JSON(200, gin.H{
		"data":       data,
		"pagination": pagination,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(201, gin.H{
		"data": data,
	})
}

func Error(c *gin.Context, status int, errType string, message string) {
	c.JSON(status, gin.H{
		"error":   errType,
		"message": message,
	})
}
