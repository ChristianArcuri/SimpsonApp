package handler

import (
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.HTML(200, "indexNew.html", gin.H{
		"title": "Phrasus - Buscador de Frases de Los Simpson",
	})
}
