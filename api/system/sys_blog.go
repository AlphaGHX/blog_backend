package system

import (
	"blog/global"
	"blog/models"
	"blog/models/request"
	"blog/models/response"
	"blog/service"
	"os"

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
		response.FailWithDetailed(err, "GetBlogInfoApi Error", c)
		return
	}
	response.OkWithData(data, c)
}

func (s *BlogApi) PostBlogFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithDetailed(err, "PostBlogFile Error", c)
		return
	}
	name := form.Value["name"]
	files := form.File["files"]

	os.Mkdir(global.CONFIG.Local.Bloghome+name[0], os.ModePerm)

	c.SaveUploadedFile(files[0], global.CONFIG.Local.Bloghome+name[0]+global.CONFIG.Local.MarkdownPath)
	c.SaveUploadedFile(files[1], global.CONFIG.Local.Bloghome+name[0]+global.CONFIG.Local.TopimgPath)

	response.Ok(c)
}

func (s *BlogApi) PostBlogInfo(c *gin.Context) {
	var data request.PostBlogInfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithDetailed(err, "PostBlogInfoApi ShouldBindJSON Error", c)
		return
	}
	insertData := models.Blog{
		Name:  data.Name,
		Title: data.Title,
		Text:  data.Text,
		Tag:   data.Tag,
	}
	err = service.ServiceGroupApp.SystemServiceGroup.BlogService.PostBlogInfo(insertData)
	if err != nil {
		response.FailWithDetailed(err, "PostBlogInfoApi service Error", c)
		return
	}
	response.Ok(c)
}
