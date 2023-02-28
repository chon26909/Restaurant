package controllers

import (
	"restaurant/models"
	"restaurant/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CustomerController interface {
	CheckIn(ctx *fiber.Ctx) error
}

type customerController struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerController(customerRepository repository.CustomerRepository) CustomerController {
	return &customerController{customerRepository}
}

func (r *customerController) CheckIn(ctx *fiber.Ctx) error {

	customer := models.Customer{
		ID:     uuid.New(),
		Qty:    1,
		Table:  1,
		Buffet: 1,
	}

	r.customerRepository.CheckIn(&customer)

	return nil
}
