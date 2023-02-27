package models

import (
	"time"

	"gorm.io/gorm"
)

type Table struct {
	ID        int `gorm:"primaryKey;"`
	Seet      int
	Available bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
