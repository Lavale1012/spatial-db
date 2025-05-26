package main

import (
	"log"

	"github.com/Lavale1012/vector-db/go-server/api"
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := "8080"
	router := gin.Default()

	group := router.Group("/api")
	{
		group.POST("/embed", api.EmbedReq)
	}

	router.Run("localhost:" + PORT)
	log.Println("Server running on http://localhost:" + PORT)
}
