package system

import (
	"blog/global"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type BlogApi struct {
}

func (s *BlogApi) GetBlog(c *gin.Context) {
	local := global.CONFIG.Local
	param := c.Param("blog")
	c.File(local.Bloghome + param + local.MarkdownPath)
}

func (s *BlogApi) GetBlogInfo(c *gin.Context) {
	param := c.Param("blog")
	data, err := service.ServiceGroupApp.SystemServiceGroup.BlogService.GetBlogInfo(param)
	if err != nil {
		c.JSON(200, gin.H{
			"info": err,
		})
	}
	c.JSON(200, data)
}
