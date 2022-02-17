package system

import (
	"blog/global"
	"blog/models/response"

	"github.com/gin-gonic/gin"
)

type ListApi struct {
}

func (s *ListApi) GetList(c *gin.Context) {
	var list response.ListResponse
	list.List = global.BLOG_LIST
	c.JSON(200, list)
}

func (s *ListApi) GetListImg(c *gin.Context) {
	param := c.Param("img")
	c.File(global.BLOG_HOME + param + "/topImg.jpg")
}
