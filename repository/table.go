package repository

import (
	"restaurant/models"

	"gorm.io/gorm"
)

type TableRepository interface {
	AddTable(table *models.Table) error
}

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepository{db}
}

func (r *tableRepository) AddTable(table *models.Table) error {
	return r.db.Create(table).Error
}
