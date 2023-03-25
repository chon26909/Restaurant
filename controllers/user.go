package controllers

import (
	"restaurant/repository"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Login(ctx *fiber.Ctx) error
}

type userController struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) UserController {
	return &userController{userRepository}
}

func (r userController) Login(ctx *fiber.Ctx) error {
	return nil
}
