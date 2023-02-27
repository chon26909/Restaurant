package models

import (
	"time"

	"gorm.io/gorm"
)

type Buffet struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
	Image       string
	Price       int16
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
