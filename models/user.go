package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Email     string
	Password  string
	Fistname  string
	Lastname  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
