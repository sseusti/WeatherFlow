package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	WriteJSON(c, http.StatusOK, gin.H{
		"status": "ok",
	})
}
