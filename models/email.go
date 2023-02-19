package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Email struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Address string
	Status  string
	UserID  uuid.UUID `gorm:"type:char(36)"`
	// User    User      `gorm:"foreignKey:UserID"`
}
