package controllers

import (
	"restaurant/dto"
	"restaurant/models"
	"restaurant/repository"

	"github.com/gofiber/fiber/v2"
)

type BuffetController interface {
	CreateBuffet(ctx *fiber.Ctx) error
	GetPackageBuffet(ctx *fiber.Ctx) error
}

type buffetController struct {
	buffetRepository repository.BuffetRepository
}

func NewBuffetController(buffetRepository repository.BuffetRepository) BuffetController {
	return &buffetController{buffetRepository}
}

func (r *buffetController) CreateBuffet(ctx *fiber.Ctx) error {

	var body dto.CreateBuffet
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	newPackage := &models.Buffet{
		Name:        body.Name,
		Description: body.Description,
		Image:       body.Image,
		Price:       body.Price,
	}

	err = r.buffetRepository.CreateBuffet(newPackage)
	if err != nil {
		return err
	}

	return nil
}

func (r *buffetController) GetPackageBuffet(ctx *fiber.Ctx) error {

	packageBuffet, err := r.buffetRepository.GetAllBuffets()
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": packageBuffet,
	})
}
