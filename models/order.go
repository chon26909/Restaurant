package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        int `gorm:"autoIncrementIncrement"`
	User      string
	Menu      int
	Table     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
