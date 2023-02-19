package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Submenu struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	MenuID      uuid.UUID `gorm:"type:char(36)"`
	Name        string
	Description string
	Image       string
	Available   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
