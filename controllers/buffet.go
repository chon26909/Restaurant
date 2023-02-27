package controllers

import (
	"restaurant/dto"
	"restaurant/models"
	"restaurant/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BuffetController interface {
	CreateBuffet(ctx *fiber.Ctx) error
	UpdateBuffet(ctx *fiber.Ctx) error
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

	newBuffet := &models.Buffet{
		Name:        body.Name,
		Description: body.Description,
		Image:       body.Image,
		Price:       body.Price,
	}

	err = r.buffetRepository.CreateBuffet(newBuffet)
	if err != nil {
		return err
	}

	return nil
}

func (r *buffetController) UpdateBuffet(ctx *fiber.Ctx) error {

	buffetID := ctx.Params("id")

	var body dto.UpdateBuffet

	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	updateBuffet := &models.Buffet{
		Name:        body.Name,
		Description: body.Description,
		Image:       body.Image,
		Price:       body.Price,
	}

	err = r.buffetRepository.UpdateBuffet(buffetID, updateBuffet)
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

	var response []dto.BuffetResponse
	for _, v := range packageBuffet {
		response = append(response, dto.BuffetResponse{
			ID:          strconv.Itoa(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Image:       v.Image,
			Price:       v.Price,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
