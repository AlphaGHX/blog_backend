package system

import (
	"blog/global"
	"blog/models"
)

type BlogService struct {
}

func (s *BlogService) GetBlogInfo(param string) (info models.Blog, err error) {
	var data models.Blog
	result := global.GROM.Where("name = ?", param).First(&data)
	return data, result.Error
}

func (s *BlogService) PostBlogInfo(data models.Blog) (err error) {
	result := global.GROM.Create(&data)
	return result.Error
}
