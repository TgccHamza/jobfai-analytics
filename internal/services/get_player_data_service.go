package services

import (
	"encoding/json"
	"fmt"
	"jobfai-analytics/internal/dto/responses"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetPlayerData retrieves the player data from the database using dbIndex and playerUUID.
func GetPlayerData(dbIndex string, playerUUID string) (map[string]interface{}, error) {
	URLMONGO := os.Getenv("URL_MONGO")
	url := fmt.Sprintf("%s/game/%s/players/%s", URLMONGO, dbIndex, playerUUID)
	client := GetHTTPClient().Client()
	resp, err := client.Get(url)
	if err != nil {
		return gin.H{}, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return gin.H{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var playerDataResponse responses.PlayerDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&playerDataResponse); err != nil {
		return gin.H{}, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	if !playerDataResponse.Success {
		return gin.H{}, fmt.Errorf("error response: %s, %s", playerDataResponse.Message, playerDataResponse.Error)
	}

	return playerDataResponse.Data, nil
}
