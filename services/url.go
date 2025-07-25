package services

import (
	"github.com/stewardyohanes/url-shortener/models"
	"github.com/stewardyohanes/url-shortener/repositories"
)

func CreateURL(url models.URL) error {
	return repositories.CreateURL(url)
}

func GetURLByShortCode(shortCode string) (models.URL, error) {
	return repositories.GetURLByShortCode(shortCode)
}

func UpdateURLVisitCount(shortCode string) error {
	return repositories.UpdateURLVisitCount(shortCode)
}
