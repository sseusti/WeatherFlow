package handler

import "github.com/gin-gonic/gin"

func WriteJSON(c *gin.Context, status int, data gin.H) {
	c.JSON(status, data)
}

func WriteError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"error": msg,
	})
}
