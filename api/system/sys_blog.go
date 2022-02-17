package system

import (
	"blog/global"

	"github.com/gin-gonic/gin"
)

type BlogApi struct {
}

func (s *ListApi) GetBlog(c *gin.Context) {
	param := c.Param("blog")
	c.File(global.BLOG_HOME + param + "/markdown.md")
}
