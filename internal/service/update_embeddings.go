package service

import (
	"log"
	"simpsonapp/config"
	"simpsonapp/internal/model"
)

// UpdateEmbeddings actualiza el campo Embedding para los registros que lo necesiten
func UpdateEmbeddings() error {
	var episodes []model.Episode

	// Obtener los registros que no tienen embeddings o tienen embedding vacío
	if err := config.DB.Where("embedding IS NULL OR embedding = ''").Find(&episodes).Error; err != nil {
		log.Printf("Error al obtener registros sin embeddings: %v", err)
		return err
	}

	log.Printf("Encontrados %d episodios para actualizar", len(episodes))

	// Recorrer los registros y actualizar los embeddings
	for _, episode := range episodes {
		log.Printf("Procesando episodio ID %d, frase: %s", episode.ID, episode.Phrase)

		embedding, err := GetEmbeddings(episode.Phrase)
		if err != nil {
			log.Printf("Error al generar embedding para frase '%s': %v", episode.Phrase, err)
			continue
		}

		log.Printf("Embedding generado correctamente, longitud: %d", len(embedding))

		// Actualizar el registro con el embedding en formato JSON
		result := config.DB.Model(&model.Episode{}).Where("id = ?", episode.ID).Update("embedding", embedding)
		if result.Error != nil {
			log.Printf("Error al guardar el embedding para ID %d: %v", episode.ID, result.Error)
		} else {
			log.Printf("Rows afectadas: %d", result.RowsAffected)
			log.Printf("Embedding actualizado para frase: '%s'", episode.Phrase)
		}
	}

	return nil
}
