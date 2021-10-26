package utils

import (
	"github.com/gin-gonic/gin"
)

func ShowErr(c *gin.Context, status int, message string, data interface{}) {
	// jika data adalah nil, maka tidak perlu membuat log untuk error ini
	if data != nil {
		CreateLog(c.FullPath(), status, message, data)
	}
	c.JSON(status, gin.H{
		"message": message,
	})
	c.Abort()
}
