package handler

import (
	"simpsonapp/config"
	"simpsonapp/internal/model"

	"github.com/gin-gonic/gin"
)

type SearchStats struct {
	TotalSearches      int64
	SuccessfulSearches int64
	TopSearches        []struct {
		Phrase string
		Count  int64
	}
}

func GetSearchStatsHandler(c *gin.Context) {
	var stats SearchStats

	// Total de búsquedas
	config.DB.Model(&model.SearchLog{}).Count(&stats.TotalSearches)

	// Búsquedas exitosas
	config.DB.Model(&model.SearchLog{}).Where("found = ?", true).Count(&stats.SuccessfulSearches)

	// Top 10 búsquedas
	config.DB.Model(&model.SearchLog{}).
		Select("search_phrase, count(*) as count").
		Group("search_phrase").
		Order("count desc").
		Limit(10).
		Scan(&stats.TopSearches)

	c.JSON(200, stats)
}
