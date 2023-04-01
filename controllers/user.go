package controllers

import (
	"fmt"
	"restaurant/dto"
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

	var body dto.User
	err := ctx.BodyParser(&body)

	fmt.Println("body", body)

	if err != nil {
		return nil
	}

	return nil
}
