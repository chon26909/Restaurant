package models

import "github.com/google/uuid"

type Menu struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Description string
	Image       string
	Submenu     []Submenu `gorm:"ForeignKey:MainMenuID"`
	Available   uint16
}
