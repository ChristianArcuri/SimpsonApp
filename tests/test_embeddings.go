package main

import (
	"fmt"
	"log"
	"simpsonapp/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	// Frase de prueba
	phrase := "Ahora va a ser una sorpresa la plancha"

	// Generar embedding
	embedding, err := service.GetEmbeddings(phrase)
	if err != nil {
		log.Println("Error al generar embedding:", err)
		return
	}

	fmt.Printf("Embedding generado: %v\n", embedding)
}
