package repository

import (
	"github.com/dhanavadh/dhanav-api/internal/models"
	"gorm.io/gorm"
)

type AboutRepository struct {
	db *gorm.DB
}

func NewAboutRepository(db *gorm.DB) *AboutRepository {
	return &AboutRepository{db: db}
}

func (r *AboutRepository) Get() (*models.About, error) {
	var about models.About
	err := r.db.First(&about).Error
	if err == gorm.ErrRecordNotFound {
		return &models.About{}, nil
	}
	if err != nil {
		return nil, err
	}
	return &about, nil
}
