package main

import (
	"log"
	"os"
	"simpsonapp/config"
	"simpsonapp/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Agregar logs de diagnóstico
	log.Printf("Iniciando aplicación...")
	log.Printf("DATABASE_URL está configurada: %v", os.Getenv("DATABASE_URL") != "")
	log.Printf("PORT está configurada: %v", os.Getenv("PORT"))
	log.Printf("GIN_MODE está configurada: %v", os.Getenv("GIN_MODE"))

	// Cargar variables de entorno solo en desarrollo
	if os.Getenv("GIN_MODE") != "release" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: No .env file found")
		}
	}

	// Conectar a la base de datos
	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// // Actualizar embeddings - comentado para evitar actualizaciones innecesarias
	// service.UpdateEmbeddings()

	// Inicializa el servidor Gin
	router := gin.Default()

	// Endpoints
	router.GET("/api/search", handler.SearchHandler)
	router.POST("/api/episodes", handler.CreateEpisodeHandler)
	router.POST("/api/update-embeddings", handler.UpdateEmbeddingsHandler)

	// Rutas para las diferentes versiones de la página
	router.GET("/", handler.HomeHandler)
	/* 	router.GET("/moe", handler.MoeHandler)         // Para index2.html
	   	router.GET("/nuclear", handler.NuclearHandler) // Para index3.html
	   	router.GET("/kwik", handler.KwikHandler)       // Para index4.html
	   	router.GET("/school", handler.SchoolHandler)   // Para index5.html
	   	router.GET("/krusty", handler.KrustyHandler)   // Para index6.html */

	// Servir archivos estáticos
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	// Inicia el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor corriendo en puerto %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
