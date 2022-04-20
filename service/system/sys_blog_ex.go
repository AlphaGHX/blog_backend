package system

import (
	"blog/global"
	"blog/models"
)

type BlogServiceEx struct{}

func (s *BlogServiceEx) PostBlogInfo(data models.Blog) (err error) {
	var oldData models.Blog
	test := global.GROM.Where("name = ?", data.Name).Find(&oldData)
	if test.RowsAffected == 0 {
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

func (s *BlogServiceEx) CreateBlogInfo(data models.Blog) (ok bool, err error) {
	var oldData models.Blog
	test := global.GROM.Where("name = ?", data.Name).Find(&oldData)
	if test.RowsAffected == 0 {
		result := global.GROM.Create(&data)
		return true, result.Error
	} else {
		return false, nil
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
