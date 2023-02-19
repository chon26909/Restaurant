package controllers

import (
	"restaurant/models"
	"restaurant/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuController interface {
	CreateMenu(ctx *fiber.Ctx) error
}

type menuController struct {
	menuRepository repository.MenuRepository
}

func NewMenuController(menuRepository repository.MenuRepository) MenuController {
	return &menuController{menuRepository}
}

func (c *menuController) CreateMenu(ctx *fiber.Ctx) error {

	body := &models.Menu{
		ID:          uuid.New(),
		Name:        "หมูสไลด์บาง",
		Description: "หมู",
		Image:       "deoaih",
		Submenus: []models.Submenu{
			{

				ID:          uuid.New(),
				Name:        "บาง",
				Description: "บางมาก",
				Image:       "sub",
				Available:   1,
			},
		},
		Available: 1,
	}

	c.menuRepository.CreateMenu(body)

	return nil
}
