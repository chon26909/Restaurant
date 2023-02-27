package repository

import (
	"restaurant/models"

	"gorm.io/gorm"
)

type BuffetRepository interface {
	CreateBuffet(data *models.Buffet) error
	GetAllBuffets() ([]models.Buffet, error)
}

type buffetRepository struct {
	db *gorm.DB
}

func NewBuffetRepository(db *gorm.DB) BuffetRepository {
	return &buffetRepository{db}
}

func (r *buffetRepository) CreateBuffet(data *models.Buffet) error {

	return r.db.Create(data).Error

}

func (r *buffetRepository) GetAllBuffets() (data []models.Buffet, err error) {

	err = r.db.Find(&data).Error

	return data, err
}
