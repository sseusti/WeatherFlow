package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Health(c *gin.Context) {
	writeJSON(c, http.StatusOK, gin.H{
		"status": "ok",
	})
}
