package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "strings"

	api "github.com/Lavale1012/vector-db/go-server/clients"
	"github.com/gin-gonic/gin"
)

type EmbedRequest struct {
	Text string `json:"text"`
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
	c.JSON(http.StatusOK, gin.H{
		"vector": vector,
	})
}
