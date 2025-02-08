package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"simpsonapp/config"
	"simpsonapp/internal/model"
	"sort"
)

// cosineSimilarity calcula la similitud del coseno entre dos vectores
func cosineSimilarity(a, b []float64) float64 {
	var dotProduct, normA, normB float64
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// SearchPhrase busca la frase más similar usando embeddings
func SearchPhrase(input string) []model.Episode {
	inputEmbeddingStr, err := GetEmbeddings(input)
	if err != nil {
		log.Printf("Error al generar embedding para la búsqueda: %v", err)
		return nil
	}

	var inputEmbedding []float64
	if err := json.Unmarshal([]byte(inputEmbeddingStr), &inputEmbedding); err != nil {
		log.Printf("Error al deserializar embedding: %v", err)
		return nil
	}

	var episodes []model.Episode
	if err := config.DB.Find(&episodes).Error; err != nil {
		log.Printf("Error al obtener episodios de la base de datos: %v", err)
		return nil
	}

	var results []model.Episode
	for _, episode := range episodes {
		if episode.Embedding == "" {
			continue
		}

		var epEmbedding []float64
		if err := json.Unmarshal([]byte(episode.Embedding), &epEmbedding); err != nil {
			continue
		}

		similarity := cosineSimilarity(inputEmbedding, epEmbedding)
		if similarity > 0.8 {
			results = append(results, episode)
		}
	}

	return results
}

func SearchSimilarPhrases(query string) ([]model.Episode, error) {
	queryEmbedding, err := GetEmbeddings(query)
	if err != nil {
		log.Printf("Error al generar embedding: %v", err)
		return nil, fmt.Errorf("error al generar embedding: %w", err)
	}

	var episodes []model.Episode
	result := config.DB.Find(&episodes)
	if result.Error != nil {
		log.Printf("Error al obtener episodios: %v", result.Error)
		return nil, result.Error
	}

	// Comparar similitud
	type episodeScore struct {
		episode model.Episode
		score   float64
	}

	var scoredEpisodes []episodeScore

	for _, ep := range episodes {
		var epEmbedding []float64
		if err := json.Unmarshal([]byte(ep.Embedding), &epEmbedding); err != nil {
			continue
		}

		var queryEmb []float64
		if err := json.Unmarshal([]byte(queryEmbedding), &queryEmb); err != nil {
			continue
		}

		score := cosineSimilarity(queryEmb, epEmbedding)
		log.Printf("Frase: '%s', Score: %f", ep.Phrase, score)

		// Bajamos el umbral a 0.5 para ser más flexible
		if score > 0.5 {
			scoredEpisodes = append(scoredEpisodes, episodeScore{ep, score})
		}
	}

	// Ordenar por similitud
	sort.Slice(scoredEpisodes, func(i, j int) bool {
		return scoredEpisodes[i].score > scoredEpisodes[j].score
	})

	// Tomar los top 5 resultados
	topResults := make([]model.Episode, 0)
	for i := 0; i < len(scoredEpisodes) && i < 5; i++ {
		topResults = append(topResults, scoredEpisodes[i].episode)
	}

	return topResults, nil
}
