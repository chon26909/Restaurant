package main

import (
	"fmt"
	"log"
	"restaurant/lib"
	"restaurant/logs"
	"restaurant/models"
	"restaurant/repository"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func main() {

	initConfig()

	db := lib.NewMySqlConnection()
	menuRepository := repository.NewMenuRepository(db)

	menuRepository.CreateMenu(&models.Menu{
		ID:          uuid.New(),
		Name:        "หมูสไลด์บาง",
		Description: "หมู",
		Image:       "",
		Submenu:     []models.Submenu{},
		Available:   0,
	})

	data, err := menuRepository.GetAllMenu()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	logs.Info("GetAllMenu")
	fmt.Println(data)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port"))))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config file %v", err))
	}
}
