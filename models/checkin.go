package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckIn struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	Table     Table     `gorm:"references:ID"`
	Buffet    Buffet    `gorm:"references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
