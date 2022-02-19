package system

import (
	"blog/global"
	"blog/models/response"
)

type TagService struct{}

func (s TagService) GetBlogUseTag(tag string) (list response.ListResponse, err error) {
	var data response.ListResponse
	result := global.GROM.Where("JSON_SEARCH(tag, 'all', ?)", tag).Find(&data.List)
	return data, result.Error
}
