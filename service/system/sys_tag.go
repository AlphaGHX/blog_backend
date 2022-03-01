package system

import (
	"blog/global"
	"blog/models"
	"blog/models/response"
)

type TagService struct{}

func (s *TagService) GetBlogUseTag(tag string) (list response.ListResponse, err error) {
	var data response.ListResponse
	var newData response.ListResponse
	result := global.GROM.Order("updated_at desc").Find(&data)
	for index, value := range data {
		for _, cTag := range value.Tag {
			if cTag == tag {
				newData = append(newData, data[index])
				break
			}
		}
	}
	return newData, result.Error
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
