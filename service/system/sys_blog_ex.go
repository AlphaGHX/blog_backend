package system

import (
	"blog/global"
	"blog/models"
	"errors"

	"gorm.io/gorm"
)

type BlogServiceEx struct{}

func (s *BlogServiceEx) PostBlogInfo(data models.Blog) (err error) {
	var oldData models.Blog
	test := global.GROM.Where("name = ?", data.Name).First(&oldData)
	if errors.Is(test.Error, gorm.ErrRecordNotFound) {
		result := global.GROM.Create(&data)
		return result.Error
	} else {
		oldData.Title = data.Title
		oldData.Text = data.Text
		oldData.Tag = data.Tag
		result := global.GROM.Save(&oldData)
		return result.Error
	}
}

func (s *BlogServiceEx) DelBlogInfo(name string) (err error) {
	getBlogInfo := new(BlogService)
	data, err := getBlogInfo.GetBlogInfo(name)
	if err != nil {
		return err
	}
	result := global.GROM.Delete(&data)
	return result.Error
}
