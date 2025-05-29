package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "strings"
	api "github.com/Lavale1012/vector-db/go-server/clients"
	"github.com/Lavale1012/vector-db/go-server/config"
	"github.com/Lavale1012/vector-db/go-server/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type EmbedRequest struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
type SearchRequest struct {
	Text string `json:"text"`
	K    int    `json:"k"`
}

func EmbedReq(c *gin.Context) {

	var req EmbedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	ReqData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	vector, err := api.EmbedClient(ReqData, c)
	if err != nil {
		fmt.Println("Error in embed_client:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get embedding vector"})
		return
	}
	model := models.VectorModel{
		ID:     req.ID,
		Text:   req.Text,
		Vector: pq.Float64Array(vector),
	}
	// data := config.DB.Create(&model).Error

	if err := config.DB.Create(&model).Error; err != nil {
		fmt.Println("Error saving vector to database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save vector to database"})
		return
	}
	fmt.Println("Vector saved to database successfully")
	c.JSON(http.StatusOK, gin.H{
		"vector": vector,
	})
}

// func SearchReq(c *gin.Context) {
// 	var req SearchRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
// 		return
// 	}
// 	ReqData, err := api.EmbedClient([]byte(req.Text), c)
// 	if err != nil {
// 		fmt.Println("Error in embed_client:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get embedding vector"})
// 		return
// 	}
// 	// for _, vector := range ReqData {}
// }
