package system

import (
	"blog/global"
	"blog/models"
	"errors"
)

type BlogService struct {
}

func (s *BlogService) GetBlogInfo(param string) (info *models.Blog, err error) {
	var data models.Blog
	result := global.GORM.Where("name = ?", param).Find(&data)
	if result.RowsAffected == 0 {
		global.STD_LOG.Errorln("BlogService.GetBlogInfo RowsAffected = 0")
		global.FILE_LOG.Errorln("BlogService.GetBlogInfo RowsAffected = 0")
		return nil, errors.New("RowsAffected = 0")
	}
	if result.Error == nil {
		data.Views++
		result = global.GORM.Save(&data)
		if result.Error != nil {
			global.STD_LOG.Errorln("BlogService.GetBlogInfo Save", result.Error.Error())
			global.FILE_LOG.Errorln("BlogService.GetBlogInfo Save", result.Error.Error())
		}
	}
	return &data, result.Error
}
