package system

import (
	"blog/service"

	"github.com/gin-gonic/gin"
)

type BlogApi struct {
}

func (s *ListApi) GetBlog(c *gin.Context) {
	param := c.Param("blog")
	c.File("./blog/" + param + "/markdown.md")
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
