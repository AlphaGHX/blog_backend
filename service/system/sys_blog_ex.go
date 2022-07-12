package system

import (
	"blog/global"
	"blog/models"
)

type BlogServiceEx struct{}

func (s *BlogServiceEx) PostBlogInfo(data models.Blog) (err error) {
	var oldData models.Blog

	test := global.GORM.Where("name = ?", data.Name).Find(&oldData)
	if test.RowsAffected == 0 {
		result := global.GORM.Create(&data)
		if result.Error != nil {
			global.LOG.Errorln("BlogServiceEx.PostBlogInfo Create", result.Error.Error())
		}

		return result.Error
	} else {
		oldData.Title = data.Title
		oldData.Text = data.Text
		oldData.Tag = data.Tag

		result := global.GORM.Save(&oldData)
		if result.Error != nil {
			global.LOG.Errorln("BlogServiceEx.PostBlogInfo Save", result.Error.Error())
		}

		return result.Error
	}
}

func (s *BlogServiceEx) CreateBlogInfo(data models.Blog) (ok bool, err error) {
	var oldData models.Blog

	test := global.GORM.Where("name = ?", data.Name).Find(&oldData)
	if test.RowsAffected == 0 {
		result := global.GORM.Create(&data)
		if result.Error != nil {
			global.LOG.Errorln("BlogServiceEx.CreateBlogInfo Create", result.Error.Error())
		}

		return true, result.Error
	} else {
		return false, nil
	}
}

func (s *BlogServiceEx) DelBlogInfo(name string) (err error) {
	getBlogInfo := new(BlogService)
	data, err := getBlogInfo.GetBlogInfo(name)
	if err != nil {
		global.LOG.Errorln("BlogServiceEx.DelBlogInfo GetBlogInfo", err.Error())
		return err
	}
	result := global.GORM.Delete(&data)
	if result.Error != nil {
		global.LOG.Errorln("BlogServiceEx.DelBlogInfo Delete", result.Error.Error())
	}
	return result.Error
}
