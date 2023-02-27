package controllers

import (
	"restaurant/dto"
	"restaurant/models"
	"restaurant/repository"

	"github.com/gofiber/fiber/v2"
)

type TableController interface {
	AddTable(ctx *fiber.Ctx) error
}

type tableController struct {
	tableRepository repository.TableRepository
}

func NewTableController(tableRepository repository.TableRepository) TableController {
	return &tableController{tableRepository}
}

func (r *tableController) AddTable(ctx *fiber.Ctx) error {

	var body dto.Table
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	newTable := &models.Table{
		ID:        body.ID,
		Seet:      body.Seet,
		Available: body.Available,
	}
	err = r.tableRepository.AddTable(newTable)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "created",
		"data":    body,
	})
}
