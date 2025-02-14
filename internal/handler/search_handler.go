package handler

import (
	"log"
	"net/http"
	"simpsonapp/config"
	"simpsonapp/internal/model"
	"simpsonapp/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

// EpisodeResponse es la estructura para la respuesta sin el campo embedding
type EpisodeResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Phrase     string    `json:"phrase"`
	Episode    string    `json:"episode"`
	Title      string    `json:"title"`
	Timestamp  string    `json:"timestamp"`
	YouTubeURL string    `json:"youtube_url"`
}

// SearchHandler maneja la búsqueda de frases
func SearchHandler(c *gin.Context) {
	phrase := c.Query("phrase")
	if phrase == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Frase de búsqueda requerida"})
		return
	}

	results, err := service.SearchSimilarPhrases(phrase)

	// Registrar la búsqueda
	searchLog := model.SearchLog{
		SearchPhrase: phrase,
		Found:        err == nil && len(results) > 0,
		ResultCount:  len(results),
		UserIP:       c.ClientIP(),
		UserAgent:    c.Request.UserAgent(),
	}

	if err := config.DB.Create(&searchLog).Error; err != nil {
		log.Printf("Error al registrar búsqueda: %v", err)
	}

	if err != nil {
		log.Printf("Error en búsqueda: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convertir los resultados al formato de respuesta
	var response []EpisodeResponse
	for _, ep := range results {
		response = append(response, EpisodeResponse{
			ID:         ep.ID,
			CreatedAt:  ep.CreatedAt,
			Phrase:     ep.Phrase,
			Episode:    ep.Episode,
			Title:      ep.Title,
			Timestamp:  ep.Timestamp,
			YouTubeURL: ep.YouTubeURL,
		})
	}

	c.JSON(http.StatusOK, response)
}

// RandomHandler maneja la solicitud para obtener una frase aleatoria
func RandomHandler(c *gin.Context) {
	result, err := service.GetRandomPhrase()
	if err != nil {
		log.Printf("Error obteniendo frase aleatoria: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := []EpisodeResponse{{
		ID:         result.ID,
		CreatedAt:  result.CreatedAt,
		Phrase:     result.Phrase,
		Episode:    result.Episode,
		Title:      result.Title,
		Timestamp:  result.Timestamp,
		YouTubeURL: result.YouTubeURL,
	}}

	c.JSON(http.StatusOK, response)
}
