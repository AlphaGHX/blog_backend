package system

import (
	"blog/models/response"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type TagApi struct{}

func (s *TagApi) GetBlogUseTag(c *gin.Context) {
	param := c.Param("tag")
	list, err := service.ServiceGroupApp.SystemServiceGroup.TagService.GetBlogUseTag(param)
	if err != nil {
		response.FailWithMessage("GetBlogUseTagApi SQL Error", c)
		return
	}
	response.OkWithData(list, c)
}

func (s *TagApi) GetTagList(c *gin.Context) {
	list, err := service.ServiceGroupApp.SystemServiceGroup.TagService.GetTagList()
	if err != nil {
		response.FailWithMessage("GetTagListApi SQL Error", c)
		return
	}
	response.OkWithData(list, c)
}
