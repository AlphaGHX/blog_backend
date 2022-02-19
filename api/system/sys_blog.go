package system

import (
	"blog/global"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type BlogApi struct {
}

func (s *ListApi) GetBlog(c *gin.Context) {
	local := global.CONFIG.Local
	param := c.Param("blog")
	c.File(local.Bloghome + param + local.MarkdownPath)
}

func (s *ListApi) GetBlogInfo(c *gin.Context) {
	param := c.Param("blog")
	data, err := service.ServiceGroupApp.SystemServiceGroup.GetBlogInfo(param)
	if err != nil {
		c.JSON(200, gin.H{
			"info": err,
		})
	}
	c.JSON(200, data)
}
