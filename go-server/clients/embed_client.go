package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type EmbedResponse struct {
	Vector []float64 `json:"vector"`
}

func EmbedClient(ReqData []byte, c *gin.Context) ([]float64, error) {
	var EmbedResp EmbedResponse
	resp, err := http.Post("http://127.0.0.1:8000/embed", "application/json", strings.NewReader(string(ReqData)))
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
