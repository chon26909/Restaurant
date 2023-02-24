package dto

import "github.com/google/uuid"

type Submenu struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}

type CreateMenuRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Submenus    []Submenu
	Available   bool `json:"available"`
}

type MenuResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Submenus  []Submenu `json:"submenus"`
	Available bool      `json:"available"`
}
