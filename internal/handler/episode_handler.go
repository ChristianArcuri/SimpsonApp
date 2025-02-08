package handler

import (
	"log"
	"net/http"
	"simpsonapp/config"
	"simpsonapp/internal/model"
	"simpsonapp/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateEpisodeHandler(c *gin.Context) {
	var episode model.Episode
	if err := c.ShouldBindJSON(&episode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Generar embedding para la frase
	embedding, err := service.GetEmbeddings(episode.Phrase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar embedding"})
		return
	}

	episode.Embedding = embedding

	// Guardar en la base de datos
	if err := config.DB.Create(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar episodio"})
		return
	}

	c.JSON(http.StatusCreated, episode)
}

func UpdateEmbeddingsHandler(c *gin.Context) {
	err := service.UpdateEmbeddings()
	if err != nil {
		log.Printf("Error actualizando embeddings: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Embeddings actualizados exitosamente"})
}
