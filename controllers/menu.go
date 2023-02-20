package controllers

import (
	"fmt"
	"restaurant/models"
	"restaurant/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuController interface {
	CreateMenu(c *fiber.Ctx) error
	GetAllMenu(c *fiber.Ctx) error
	UpdateMenu(c *fiber.Ctx) error
}

type menuController struct {
	menuRepository repository.MenuRepository
}

func NewMenuController(menuRepository repository.MenuRepository) MenuController {
	return &menuController{menuRepository}
}

func (r *menuController) CreateMenu(c *fiber.Ctx) error {

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

	r.menuRepository.CreateMenu(body)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "created",
	})
}

func (r *menuController) GetAllMenu(c *fiber.Ctx) error {
	allMenu, err := r.menuRepository.GetAllMenu()
	if err != nil {
		return err
	}

	fmt.Printf("all menu %v", allMenu)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    allMenu,
	})
}

func (r *menuController) UpdateMenu(c *fiber.Ctx) error {
	return nil
}
