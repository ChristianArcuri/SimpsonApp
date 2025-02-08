package handler

import (
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Buscador de Frases de Los Simpson",
	})
}
