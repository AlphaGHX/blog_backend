package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Blog struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Name      string    `gorm:"unique" json:"name"`
	Title     string    `gorm:"not null" json:"title"`
	Text      string    `gorm:"not null" json:"text"`
	Tag       Tag       `gorm:"type:json" json:"tag"`
}

type Tag []string

func (c Tag) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Tag) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
