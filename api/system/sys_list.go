package system

import (
	"blog/global"
	"blog/models/response"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type ListApi struct {
}

func (s *ListApi) GetList(c *gin.Context) {
	list, err := service.ServiceGroupApp.SystemServiceGroup.ListService.GetBlogList()
	if err != nil {
		response.FailWithMessage("GetListApi SQL Error", c)
		return
	}
	response.OkWithData(list, c)
}

func (s *ListApi) GetListImg(c *gin.Context) {
	local := global.CONFIG.Local
	param := c.Param("img")
	c.File(local.Bloghome + param + local.TopimgPath)
}
