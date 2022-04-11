package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Blog struct {
	Name      string    `gorm:"primarykey;not null;unique" json:"name"`
	Title     string    `gorm:"not null" json:"title"`
	Text      string    `gorm:"not null" json:"text"`
	Views     uint      `gorm:"not null;default:0" json:"views"`
	Tag       Tag       `gorm:"type:string" json:"tag"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type Tag []string

// 序列化
func (c Tag) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

// 反序列化
func (c *Tag) Scan(input interface{}) error {
	return json.Unmarshal([]byte(input.(string)), c)
}
