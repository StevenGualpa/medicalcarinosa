package repository

import (
	"GolandProyectos/models"
	"gorm.io/gorm"
)

type DeviceTokenRepository interface {
	GetAllDeviceTokens() ([]models.DeviceToken, error)
}

type deviceTokenRepository struct {
	db *gorm.DB
}

func NewDeviceTokenRepository(db *gorm.DB) DeviceTokenRepository {
	return &deviceTokenRepository{db}
}

func (r *deviceTokenRepository) GetAllDeviceTokens() ([]models.DeviceToken, error) {
	var tokens []models.DeviceToken
	result := r.db.Find(&tokens)
	return tokens, result.Error
}
