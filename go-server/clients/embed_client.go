package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type EmbedResponse struct {
	Vector []float64 `json:"vector"`
}

func EmbedClient(ReqData []byte, c *gin.Context) ([]float64, error) {
	var EmbedResp EmbedResponse
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load environment variables"})
		return nil, err
	}
	EmbedUrl := os.Getenv("EMBED_API_URL")
	if EmbedUrl == "" {
		fmt.Println("EMBED_URL not set in .env file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Embedding service URL not configured"})
		return nil, fmt.Errorf("embedding service URL not configured")
	}
	resp, err := http.Post(EmbedUrl, "application/json", strings.NewReader(string(ReqData)))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request to embedding service"})
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&EmbedResp)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response from embedding service"})
		return nil, err
	}
	return EmbedResp.Vector, nil
}
