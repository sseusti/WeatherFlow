package handler

import "github.com/gin-gonic/gin"

func writeJSON(c *gin.Context, status int, data gin.H) {
	c.JSON(status, data)
}

func writeError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"error": msg,
	})
}
