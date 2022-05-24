package system

import (
	"blog/global"
	"blog/models/response"
)

type ListService struct {
}

func (s *ListService) GetBlogList() (BlogList *response.ListResponse, err error) {
	var data response.ListResponse

	result := global.GROM.Order("created_at desc").Find(&data)
	if result.Error != nil {
		global.STD_LOG.Errorln("ListService.GetBlogList Find", result.Error.Error())
		global.FILE_LOG.Errorln("ListService.GetBlogList Find", result.Error.Error())
		return nil, result.Error
	}

	return &data, result.Error
}
