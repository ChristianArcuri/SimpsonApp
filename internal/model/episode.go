package model

import "gorm.io/gorm"

type Episode struct {
	gorm.Model
	Phrase     string `gorm:"index"`
	Episode    string
	Title      string
	Timestamp  string
	YouTubeURL string
	Embedding  string `gorm:"type:text"`
}
