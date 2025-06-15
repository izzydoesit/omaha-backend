package services

import (
	"github.com/izzydoesit/omaha-backend/internal/models"
	"gorm.io/gorm"
)

func SaveHand(db *gorm.DB, hand *models.Hand) error {
	return db.Create(hand).Error
}

func GetHandsByUser(db *gorm.DB, userID string) ([]models.Hand, error) {
	var hands []models.Hand
	err := db.Where("user_id = ?", userID).Order("created_at desc").Find(&hands).Error
	return hands, err
}