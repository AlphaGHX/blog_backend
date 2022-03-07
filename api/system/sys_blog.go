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
		response.FailWithMessage("GetBlogInfoApi SQL Error", c)
		return
	}
	response.OkWithData(data, c)
}

func (s *BlogApi) PostBlogFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage("PostBlogFile Recv MultipartForm Error", c)
		return
	}
	name := form.Value["name"]
	markdown := form.File["markdown"]
	topImg := form.File["topImg"]

	os.Mkdir(global.CONFIG.Local.Bloghome+name[0], os.ModePerm)

	if len(markdown) > 0 {
		c.SaveUploadedFile(markdown[0], global.CONFIG.Local.Bloghome+name[0]+global.CONFIG.Local.MarkdownPath)
	}

	if len(topImg) > 0 {
		c.SaveUploadedFile(topImg[0], global.CONFIG.Local.Bloghome+name[0]+global.CONFIG.Local.TopimgPath)
	}

	response.Ok(c)
}

func (s *BlogApi) PostBlogInfo(c *gin.Context) {
	var data request.PostBlogInfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage("PostBlogInfoApi ShouldBindJSON Error", c)
		return
	}
	insertData := models.Blog{
		Name:  data.Name,
		Title: data.Title,
		Text:  data.Text,
		Tag:   data.Tag,
	}
	err = service.ServiceGroupApp.SystemServiceGroup.BlogServiceEx.PostBlogInfo(insertData)
	if err != nil {
		response.FailWithMessage("PostBlogInfoApi SQL Error", c)
		return
	}
	response.Ok(c)
}

func (s *BlogApi) CreateBlogInfo(c *gin.Context) {
	var data request.PostBlogInfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage("PostBlogInfoApi ShouldBindJSON Error", c)
		return
	}
	insertData := models.Blog{
		Name:  data.Name,
		Title: data.Title,
		Text:  data.Text,
		Tag:   data.Tag,
	}
	var ok bool
	ok, err = service.ServiceGroupApp.SystemServiceGroup.BlogServiceEx.CreateBlogInfo(insertData)
	if err != nil {
		response.FailWithMessage("PostBlogInfoApi SQL Error", c)
		return
	}
	if ok {
		response.Ok(c)
	} else {
		response.FailWithMessage("PostBlogInfoApi name is exist Error", c)
	}
}

func (s *BlogApi) DelBlogInfo(c *gin.Context) {
	param := c.Param("blog")
	err := service.ServiceGroupApp.SystemServiceGroup.BlogServiceEx.DelBlogInfo(param)
	if err != nil {
		response.FailWithMessage("DelBlogInfoApi SQL Error", c)
	}
	err = os.RemoveAll(global.CONFIG.Local.Bloghome + param)
	if err != nil {
		response.FailWithMessage("DelBlogInfoApi file operations Error", c)
	}
	response.Ok(c)
}
