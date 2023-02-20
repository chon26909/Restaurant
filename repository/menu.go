package repository

import (
	"restaurant/models"

	"gorm.io/gorm"
)

type MenuRepository interface {
	CreateMenu(menu *models.Menu) error
	GetAllMenu() (menu []models.Menu, err error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db}
}

func (r *menuRepository) GetAllMenu() (menu []models.Menu, err error) {

	// Submenus name in model Menu
	err = r.db.Preload("Submenus").Find(&menu).Error

	return menu, err
}

func (r *menuRepository) CreateMenu(menu *models.Menu) error {
	return r.db.Create(menu).Error
}
