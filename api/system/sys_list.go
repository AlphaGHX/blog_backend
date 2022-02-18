package system

import (
	"blog/service"

	"github.com/gin-gonic/gin"
)

type ListApi struct {
}

func (s *ListApi) GetList(c *gin.Context) {
	list, err := service.ServiceGroupApp.SystemServiceGroup.GetBlogList()
	if err != nil {
		c.JSON(200, gin.H{
			"info": err,
		})
	}
	c.JSON(200, list)
}

func (s *ListApi) GetListImg(c *gin.Context) {
	param := c.Param("img")
	c.File("./blog/" + param + "/topImg.jpg")
}
