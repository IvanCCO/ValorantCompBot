package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ValorantAgent struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"displayName"`
}

type ValorantResponse struct {
	Data []ValorantAgent `json:"data"`
}

var URI string

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	res := os.Getenv("VALORANT_API_PATH")
	if res == "" {
		log.Fatal("VALORANT_API_PATH not found in .env file")
	}

	URI = res
}

func GetAgents() ([]ValorantAgent, error) {
	resp, err := http.Get(URI)
	if err != nil {
		return nil, errors.New("failed to retrieve Valorant agents")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var response ValorantResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("error unmarshalling JSON")
	}

	return response.Data, nil
}
