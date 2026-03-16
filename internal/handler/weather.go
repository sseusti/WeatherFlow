package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCurrentWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		writeError(c, http.StatusBadRequest, "city is required")
		return
	}

	response, err := h.weatherService.GetCurrent(c.Request.Context(), city)
	if err != nil {
		log.Printf("failed to fetch external weather status: %v", err)
		writeError(c, http.StatusInternalServerError, "failed to fetch external weather status")
		return
	}

	writeJSON(c, http.StatusOK, response)
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

func (h *Handler) GetHourlyWeatherByCity(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		writeError(c, http.StatusBadRequest, "city is required")
		return
	}

	response, err := h.weatherService.GetHourly(c.Request.Context(), city)
	if err != nil {
		log.Printf("failed to fetch external weather status: %v", err)
		writeError(c, http.StatusInternalServerError, "failed to fetch external weather status")
		return
	}

	writeJSON(c, http.StatusOK, response)
	return
}
