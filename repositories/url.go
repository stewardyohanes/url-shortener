package repositories

import (
	"github.com/stewardyohanes/url-shortener/config"
	"github.com/stewardyohanes/url-shortener/models"
	"gorm.io/gorm"
)

func CreateURL(url models.URL) error {
	return config.DB.Create(&url).Error
}

func GetURLByShortCode(shortCode string) (models.URL, error) {
	var url models.URL
	if err := config.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return models.URL{}, err
	}
	return url, nil
}

func UpdateURLVisitCount(shortCode string) error {
	return config.DB.Model(&models.URL{}).Where("short_code = ?", shortCode).Update("visit_count", gorm.Expr("visit_count + 1")).Error
}
