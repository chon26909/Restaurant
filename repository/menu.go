package repository

import (
	"restaurant/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	CreateMenu(menu *models.Menu) error
	GetAllMenu() (menu []models.Menu, err error)
	UpdateMenu(menuID uuid.UUID, menu *models.Menu) error
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

func (r *menuRepository) UpdateMenu(menuID uuid.UUID, menu *models.Menu) error {

	r.db.First(&menu)
	return r.db.Save(&menu).Error
}
