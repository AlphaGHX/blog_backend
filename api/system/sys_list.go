package system

import (
	"blog/global"
	"blog/models/response"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type ListApi struct {
}

// @Summary 获取含有所有博客信息的列表
// @Router /system/list [get]
// @Verify none
// @Param none
// @Success 200 createdat: "创建时间" updatedat: "更新时间" name: "博客名字" title: "博客标题" text: "博客简介" views: "浏览量" tag: [..."标签"]
func (s *ListApi) GetList(c *gin.Context) {
	list, err := service.ServiceGroupApp.SystemServiceGroup.ListService.GetBlogList()
	if err != nil {
		response.FailWithMessage("GetListApi SQL Error", c)
		return
	}
	response.OkWithData(list, c)
}

// @Summary 通过博客名字获取对应的图片
// @Router /system/list/img/:img [get]
// @Verify none
// @Param uri "博客名字"
// @Success 200 "博客顶部图片"
func (s *ListApi) GetListImg(c *gin.Context) {
	local := global.CONFIG.Local
	param := c.Param("img")
	c.File(local.Bloghome + param + local.TopimgPath)
}
