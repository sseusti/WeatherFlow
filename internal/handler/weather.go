package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCurrentWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "city is required",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"city":        city,
			"temperature": 0,
			"condition":   "stub",
		},
	)

}
