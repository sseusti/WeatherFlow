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

	writeJSON(c, http.StatusOK, h.weatherService.GetCurrent(city))
	return
}

func (h *Handler) GetWeatherByCity(c *gin.Context) {
	city := c.Param("city")
	if city == "" {
		writeError(c, http.StatusBadRequest, "city is required")
		return
	}

	writeJSON(c, http.StatusOK, h.weatherService.GetByCity(city))
	return
}
