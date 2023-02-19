package models

import "github.com/google/uuid"

type Submenu struct {
	ID          uuid.UUID `gorm:"primaryKey;column:id"`
	MainMenuID  uuid.UUID
	Name        string
	Description string
	Image       string
	Available   int
}
