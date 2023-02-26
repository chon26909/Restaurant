package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Package struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	Name        string
	Description string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
