package repository

import (
	"restaurant/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CheckIn(customer *models.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) CheckIn(customer *models.Customer) error {

	return r.db.Create(customer).Error
}
