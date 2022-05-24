package system

import (
	"blog/global"
	"blog/models"
	"blog/models/response"
)

type TagService struct{}

func (s *TagService) GetBlogUseTag(tag string) (list *response.ListResponse, err error) {
	var data response.ListResponse
	var newData response.ListResponse
	result := global.GROM.Order("created_at desc").Find(&data)
	if result.Error != nil {
		global.STD_LOG.Errorln("TagService.GetBlogUseTag Find", result.Error.Error())
		global.FILE_LOG.Errorln("TagService.GetBlogUseTag Find", result.Error.Error())
		return nil, result.Error
	}
	for index, value := range data {
		for _, cTag := range value.Tag {
			if cTag == tag {
				newData = append(newData, data[index])
				break
			}
		}
	}
	return &newData, result.Error
}

func (s *TagService) GetTagList() (list *models.Tag, err error) {
	var tempData response.ListResponse
	var data models.Tag
	result := global.GROM.Select("tag").Find(&tempData)
	if result.Error != nil {
		global.STD_LOG.Errorln("TagService.GetTagList Find", result.Error.Error())
		global.FILE_LOG.Errorln("TagService.GetTagList Find", result.Error.Error())
		return nil, result.Error
	}
	for _, v := range tempData {
		data = append(data, v.Tag...)
	}
	return &data, result.Error
}
