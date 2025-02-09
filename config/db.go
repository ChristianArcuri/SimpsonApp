package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB es la conexión a la base de datos
var DB *gorm.DB

func ConnectDatabase() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL no está configurada")
	}
	// Agregar sslmode=require para Railway
	if os.Getenv("RAILWAY_ENVIRONMENT") != "" {
		dbURL += "?sslmode=require"
	}
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	DB = database
	return nil
}
