package system

import (
	"blog/models/response"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type TagApi struct{}

func (*TagApi) GetBlogUseTag(c *gin.Context) {
	param := c.Param("tag")
	list, err := service.ServiceGroupApp.SystemServiceGroup.TagService.GetBlogUseTag(param)
	if err != nil {
		response.FailWithDetailed(err, "GetBlogUseTagApi Error", c)
		return
	}
	response.OkWithData(list, c)
}
