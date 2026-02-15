package main

import (
	"pricebite/backend/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	api.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}

