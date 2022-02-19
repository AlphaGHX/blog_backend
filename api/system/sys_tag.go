package system

import (
	"blog/service"

	"github.com/gin-gonic/gin"
)

type TagApi struct{}

func (*TagApi) GetBlogUseTag(c *gin.Context) {
	param := c.Param("tag")
	list, err := service.ServiceGroupApp.SystemServiceGroup.TagService.GetBlogUseTag(param)
	if err != nil {
		c.JSON(200, gin.H{
			"info": err,
		})
	}
	c.JSON(200, list)
}
