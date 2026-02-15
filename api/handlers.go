package api

import (
	"net/http"
	"strings"

	"pricebite/backend/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", healthHandler)
	r.GET("/apple", appleMultiHandler)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Price Bite backend running",
	})
}

func appleMultiHandler(c *gin.Context) {
	countryParam := c.Query("countries")

	if countryParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "countries query param required",
		})
		return
	}

	countries := strings.Split(countryParam, ",")

	data, err := service.FetchAllCountries(countries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

