package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCurrentWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		writeError(c, http.StatusBadRequest, "city is required")
		return
	}

	writeJSON(c, http.StatusOK, gin.H{
		"city":        city,
		"temperature": 0,
		"condition":   "stub",
	})
	return
}

func (h *Handler) GetWeatherByCity(c *gin.Context) {
	city := c.Param("city")
	if city == "" {
		writeError(c, http.StatusBadRequest, "city is required")
		return
	}

	writeJSON(c, http.StatusOK, gin.H{
		"city":        city,
		"source":      "path",
		"temperature": 0,
		"condition":   "stub",
	})
	return
}
