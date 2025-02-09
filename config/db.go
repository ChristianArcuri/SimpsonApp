package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB es la conexión a la base de datos
var DB *gorm.DB

func ConnectDatabase() error {
	dbURL := os.Getenv("DATABASE_URL")
	log.Printf("Intentando conectar a la base de datos...")

	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL no está configurada")
	}

	log.Printf("URL de base de datos encontrada")

	// Siempre usar SSL para conexiones externas
	if !strings.Contains(dbURL, "sslmode=") {
		dbURL += "?sslmode=require"
	}

	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	DB = database
	log.Printf("Conexión exitosa a la base de datos")
	return nil
}
