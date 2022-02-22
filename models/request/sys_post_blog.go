package request

import "blog/models"

type PostBlogInfo struct {
	Name  string     `json:"name" binding:"required"`
	Title string     `json:"title" binding:"required"`
	Text  string     `json:"text" binding:"required"`
	Tag   models.Tag `json:"tag" binding:"required"`
}
