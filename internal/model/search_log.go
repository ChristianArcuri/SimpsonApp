package model

import "gorm.io/gorm"

type SearchLog struct {
	gorm.Model
	SearchPhrase string `gorm:"index"`
	Found        bool   // true si se encontraron resultados
	ResultCount  int    // número de resultados encontrados
	UserIP       string // opcional: IP del usuario
	UserAgent    string // opcional: navegador/dispositivo
}
