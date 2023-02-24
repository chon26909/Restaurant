package repository

import (
	"fmt"
	"restaurant/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	CreateMenu(menu *models.Menu) error
	GetAllMenu() (menu []models.Menu, err error)
	UpdateMenu(menuID uuid.UUID, menu *models.Menu) error
	DeleteMenu(menuID uuid.UUID) error
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

	// var oldData model.
	err := r.db.First(&menu)
	if err != nil {
		fmt.Println("")
		fmt.Printf("error => %v\n", err)
	}

	fmt.Printf("success %v\n", menu)

	return r.db.Save(&menu).Error

	// return r.db.Model(&menu).Where("id = ?", menuID).Updates(menu).Error
}

func (r *menuRepository) DeleteMenu(menuID uuid.UUID) error {

	return r.db.Delete(&models.Menu{}, menuID).Error
}
