package main

import (
	"log"

	"github.com/Lavale1012/vector-db/go-server/api"
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := "8080"
	router := gin.Default()

	router.Use(gin.Recovery()) // Use recovery middleware to handle panics
	router.Use(gin.Logger())   // Use logger middleware to log requests
	api.ApiRoutes(router)

	if err := router.Run("localhost:" + PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server running on http://localhost:" + PORT)
}
