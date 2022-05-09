package system

import (
	"blog/global"
	"blog/models"
	"errors"
)

type BlogService struct {
}

func (s *BlogService) GetBlogInfo(param string) (info models.Blog, err error) {
	var data models.Blog
	result := global.GROM.Where("name = ?", param).Find(&data)
	if result.RowsAffected == 0 {
		return data, errors.New("nullRows")
	}
	if result.Error == nil {
		data.Views++
		global.GROM.Save(&data)
	}
	return data, result.Error
}
