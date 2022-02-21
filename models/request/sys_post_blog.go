package request

import "blog/models"

type PostBlogInfo struct {
	Name  string     `gorm:"unique" json:"name" binding:"required"`
	Title string     `gorm:"not null" json:"title" binding:"required"`
	Text  string     `gorm:"not null" json:"text" binding:"required"`
	Tag   models.Tag `gorm:"type:json" json:"tag" binding:"required"`
}
