package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCurrentWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		WriteError(c, http.StatusBadRequest, "city is required")
		return
	}

	WriteJSON(c, http.StatusOK, gin.H{
		"city":        city,
		"temperature": 0,
		"condition":   "stub",
	})
	return
}

func GetWeatherByCity(c *gin.Context) {
	city := c.Param("city")
	if city == "" {
		WriteError(c, http.StatusBadRequest, "city is required")
		return
	}

	WriteJSON(c, http.StatusOK, gin.H{
		"city":        city,
		"source":      "path",
		"temperature": 0,
		"condition":   "stub",
	})
	return
}
