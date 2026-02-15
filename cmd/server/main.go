package main

import (
	"log"
	"os"
	"pricebite/backend/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	
	corsOrigin := os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "*" 
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	
	api.RegisterRoutes(router)

	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}
	log.Printf("Server running on port %s", port)

	if err := router.Run(":" + port); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
