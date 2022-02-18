package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Title string `gorm:"not null"`
	Text  string `gorm:"not null"`
}
