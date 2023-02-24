package controllers

import (
	"fmt"
	"restaurant/dto"
	"restaurant/models"
	"restaurant/repository"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuController interface {
	CreateMenu(ctx *fiber.Ctx) error
	GetAllMenu(ctx *fiber.Ctx) error
	UpdateMenu(ctx *fiber.Ctx) error
	DeleteMenu(ctx *fiber.Ctx) error
}

type menuController struct {
	menuRepository repository.MenuRepository
}

func NewMenuController(menuRepository repository.MenuRepository) MenuController {
	return &menuController{menuRepository}
}

func (r *menuController) CreateMenu(ctx *fiber.Ctx) error {

	var body dto.CreateMenuRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	// fmt.Printf("body %v", body)

	resultUpload, err := utils.UploadImage(body.Image)
	if err != nil {
		return err
	}

	fmt.Printf("%v", resultUpload)

	menuID := uuid.New()

	var subMenus []models.Submenu
	for _, v := range body.Submenus {
		subMenus = append(subMenus, models.Submenu{
			ID:          uuid.New(),
			MenuID:      menuID,
			Name:        v.Name,
			Description: v.Description,
			Available:   v.Available,
		})
	}

	newMenu := &models.Menu{
		ID:          menuID,
		Name:        body.Name,
		Description: body.Description,
		Image:       utils.GetImageNameFromUrl(resultUpload.SecureURL),
		Submenus:    subMenus,
		Available:   true,
	}

	err = r.menuRepository.CreateMenu(newMenu)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "created",
	})

}

func (r *menuController) GetAllMenu(ctx *fiber.Ctx) error {
	allMenu, err := r.menuRepository.GetAllMenu()
	if err != nil {
		return err
	}

	var response []dto.MenuResponse
	for _, v := range allMenu {

		var subMenus []dto.Submenu

		for _, s := range v.Submenus {
			subMenus = append(subMenus, dto.Submenu{
				Name:        s.Name,
				Description: s.Description,
				Available:   s.Available,
			})
		}

		response = append(response, dto.MenuResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Image:       utils.GetPublicUrl(v.Image),
			Submenus:    subMenus,
			Available:   v.Available,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func (r *menuController) UpdateMenu(ctx *fiber.Ctx) error {

	// menuID := uuid.MustParse(ctx.Params("id"))

	// body := &models.Menu{
	// 	ID:          uuid.MustParse("d8630e24-2e8e-4d4f-82f7-296d4f6f19e1"),
	// 	Name:        "หมูสไลด์",
	// 	Description: "เหมาะสำหรับชาบู",
	// 	Image:       "deoaih",
	// 	Submenus: []models.Submenu{
	// 		{
	// 			ID:          uuid.MustParse("937199db-929d-4f24-b408-de7af66b602b"),
	// 			MenuID:      uuid.MustParse("d8630e24-2e8e-4d4f-82f7-296d4f6f19e1"),
	// 			Name:        "บาง",
	// 			Description: "บางมากกกกกกกกกกกก",
	// 			Image:       "sub",
	// 			Available:   1,
	// 		},
	// 	},
	// 	Available: 1,
	// }

	// err := r.menuRepository.UpdateMenu(menuID, body)
	// if err != nil {
	// 	return err
	// }

	// return err

	return nil
}

func (r *menuController) DeleteMenu(ctx *fiber.Ctx) error {

	menuID := uuid.MustParse(ctx.Params("id"))

	err := r.menuRepository.DeleteMenu(menuID)
	if err != nil {
		return err
	}

	return nil
}
