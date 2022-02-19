package models

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Title string `gorm:"not null"`
	Text  string `gorm:"not null"`
	Tag   Tag    `gorm:"type:json"`
}

type Tag []string

func (c Tag) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Tag) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
