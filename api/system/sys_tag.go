package system

import (
	"blog/models/response"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type TagApi struct{}

// @Summary 通过标签获得对应的博客
// @Router /system/tag/:tag [get]
// @Verify none
// @Param uri "标签"
// @Success 200 createdat: "创建时间" updatedat: "更新时间" name: "博客名字" title: "博客标题" text: "博客简介" views: "浏览量" tag: [..."标签"]
func (s *TagApi) GetBlogUseTag(c *gin.Context) {
	param := c.Param("tag")
	list, err := service.ServiceGroupApp.SystemServiceGroup.TagService.GetBlogUseTag(param)
	if err != nil {
		response.FailWithMessage("GetBlogUseTagApi SQL Error", c)
		return
	}
	response.OkWithData(list, c)
}

// @Summary 获取所有标签的列表
// @Router /system/tag [get]
// @Verify none
// @Param none
// @Success 200 [..."标签"]
func (s *TagApi) GetTagList(c *gin.Context) {
	list, err := service.ServiceGroupApp.SystemServiceGroup.TagService.GetTagList()
	if err != nil {
		response.FailWithMessage("GetTagListApi SQL Error", c)
		return
	}
	response.OkWithData(list, c)
}
