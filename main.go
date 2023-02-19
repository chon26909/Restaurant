package main

import (
	"fmt"
	"log"
	"restaurant/controllers"
	"restaurant/lib"
	"restaurant/models"
	"restaurant/repository"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Test()

	// return

	initConfig()

	db := lib.NewMySqlConnection()
	// dsn := "restaurant:P@ssw0rd@tcp(127.0.0.1:3306)/restaurant?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: &lib.SqlLogger{},
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	// err = db.AutoMigrate(&models.Menu{}, &models.Submenu{})
	// if err != nil {
	// 	log.Fatalf("Failed to auto-migrate models: %v", err)
	// }
	menuRepository := repository.NewMenuRepository(db)

	menuController := controllers.NewMenuController(menuRepository)

	// menu := &models.Menu{
	// 	ID:          uuid.New(),
	// 	Name:        "หมูสไลด์บาง",
	// 	Description: "หมู",
	// 	Image:       "deoaih",
	// 	Submenus: []models.Submenu{
	// 		{

	// 			ID:          uuid.New(),
	// 			Name:        "บาง",
	// 			Description: "บางมาก",
	// 			Image:       "sub",
	// 			Available:   1,
	// 		},
	// 	},
	// 	Available: 1,
	// }

	// menuRepository.CreateMenu(menu)

	// data, err := menuRepository.GetAllMenu()
	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }
	// logs.Info("GetAllMenu")
	// fmt.Println(data)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	menu := app.Group("/menu")
	menu.Post("/", menuController.CreateMenu)

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

func Test() {

	// Connect to the MySQL database
	dsn := "restaurant:P@ssw0rd@tcp(127.0.0.1:3306)/restaurant?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &lib.SqlLogger{},
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the models
	err = db.AutoMigrate(&models.User{}, &models.Email{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	// Create a sample user and emails
	user := models.User{
		ID:   uuid.New(),
		Name: "John Doe",
		Emails: []models.Email{
			{
				ID:      uuid.New(),
				Address: "john.doe@example.com",
				Status:  "false",
			},
			{
				ID:      uuid.New(),
				Address: "johndoe@gmail.com",
				Status:  "false",
			},
		},
	}

	// Insert the user and emails into the database
	err = db.Create(&user).Error
	if err != nil {
		log.Fatalf("Failed to insert user and emails: %v", err)
	}

	// Print the user and email data from the database
	var userFromDB models.User
	err = db.Preload("Emails").First(&userFromDB, "name = ?", "John Doe").Error
	if err != nil {
		log.Fatalf("Failed to get user from database: %v", err)
	}
	fmt.Printf("User: %v\n", userFromDB)
	for _, email := range userFromDB.Emails {
		fmt.Printf("Email: %v\n", email)
	}
}
