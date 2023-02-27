package repository

import "gorm.io/gorm"

type ReceptionRepository interface {
	CheckIn() error
}

type receptionRepository struct {
	db *gorm.DB
}

func NewReceptionRepository(db *gorm.DB) ReceptionRepository {
	return &receptionRepository{db}
}

func (r *receptionRepository) CheckIn() error {
	return nil
}
