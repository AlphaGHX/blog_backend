package system

import (
	"blog/global"
	"blog/models"
	"blog/models/response"
)

type TagService struct{}

func (s *TagService) GetBlogUseTag(tag string) (list response.ListResponse, err error) {
	var data response.ListResponse
	result := global.GROM.Where("JSON_SEARCH(tag, 'all', ?)", tag).Find(&data)
	return data, result.Error
}

func (s *TagService) GetTagList() (list models.Tag, err error) {
	var tempData response.ListResponse
	var data models.Tag
	result := global.GROM.Select("tag").Find(&tempData)
	for _, v := range tempData {
		data = append(data, v.Tag...)
	}
	return data, result.Error
}
