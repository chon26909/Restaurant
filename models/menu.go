package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name        string
	Description string
	Image       string
	Submenus    []Submenu
	Available   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
