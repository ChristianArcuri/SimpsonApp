package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// GetEmbeddings genera un embedding para un texto usando OpenAI
func GetEmbeddings(text string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API Key no configurada")
	}

	client := resty.New()

	requestBody := map[string]interface{}{
		"input": text,
		"model": "text-embedding-3-small",
	}

	var response struct {
		Data []struct {
			Embedding []float64 `json:"embedding"`
		} `json:"data"`
	}

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		SetResult(&response).
		Post("https://api.openai.com/v1/embeddings")

	if err != nil || resp.StatusCode() != 200 {
		return "", fmt.Errorf("error en la API de OpenAI: %v", err)
	}

	if len(response.Data) == 0 {
		return "", fmt.Errorf("no se generaron embeddings")
	}

	jsonEmbedding, err := json.Marshal(response.Data[0].Embedding)
	if err != nil {
		return "", fmt.Errorf("error al convertir embedding a JSON: %w", err)
	}

	return string(jsonEmbedding), nil
}
