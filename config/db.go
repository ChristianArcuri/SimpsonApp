package config

import (
	"fmt"
	"log"
	"os"
	"simpsonapp/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB es la conexión a la base de datos
var DB *gorm.DB

func ConnectDatabase() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Fallback a configuración local
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	}

	// Abre la conexión
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	// Migrar modelos
	if err = database.AutoMigrate(&model.Episode{}); err != nil {
		return fmt.Errorf("error en la migración: %w", err)
	}

	// Asegurarnos que el campo embedding sea de tipo text
	if err = database.Exec("ALTER TABLE episodes ALTER COLUMN embedding TYPE text").Error; err != nil {
		log.Printf("Warning al modificar tipo de embedding: %v", err)
	}

	log.Println("Conexión exitosa a PostgreSQL")
	DB = database
	return nil
}
