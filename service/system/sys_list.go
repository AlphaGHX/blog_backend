package system

import (
	"blog/global"
	"blog/models/response"
)

type ListService struct {
}

func (s *ListService) GetBlogList() (BlogList response.ListResponse, err error) {
	var data response.ListResponse
	result := global.GROM.Find(&data.List)
	return data, result.Error
}
