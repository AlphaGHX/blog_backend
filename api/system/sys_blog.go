package system

import (
	"blog/models"
	"blog/models/request"
	"blog/models/response"
	"blog/service"
	"blog/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type BlogApi struct {
}

// @Summary 通过博客名称获取博客的markdown文档
// @Router /system/blog/:blog [get]
// @Verify none
// @Param uri 博客名称
// @Success 200 text 博客内容
func (s *BlogApi) GetBlog(c *gin.Context) {
	param := c.Param("blog")
	c.File(utils.GetBlogMarkdownPath(param))
}

// @Summary 通过博客名称获取博客的详细信息
// @Router /system/blog-info/:blog [get]
// @Verify none
// @Param uri 博客名称
// @Success 200 json 博客详细信息
func (s *BlogApi) GetBlogInfo(c *gin.Context) {
	param := c.Param("blog")
	data, err := service.ServiceGroupApp.SystemServiceGroup.BlogService.GetBlogInfo(param)
	if err != nil {
		response.FailWithMessage("GetBlogInfoApi SQL Error", c)
		return
	}
	response.OkWithData(data, c)
}

// @Summary 创建或更新博客的图片和markdown文档
// @Router /system/blog/uploadfile [post]
// @Verify JWT
// @Param headers x-token "JWT" body form-data key:{files: [markdownFile, imgFile], name: "博客名字"}
// @Success 200
func (s *BlogApi) PostBlogFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage("PostBlogFile Recv MultipartForm Error", c)
		return
	}
	name := form.Value["name"]
	markdown := form.File["markdown"]
	topImg := form.File["topImg"]

	os.Mkdir(utils.GetBlogPath(name[0]), os.ModePerm)

	if len(markdown) > 0 {
		c.SaveUploadedFile(markdown[0], utils.GetBlogMarkdownPath(name[0]))
	}

	if len(topImg) > 0 {
		c.SaveUploadedFile(topImg[0], utils.GetBlogTopimgPath(name[0]))
	}

	response.Ok(c)
}

// @Summary 更新博客的详细信息
// @Router /system/blog/uploadinfo [post]
// @Verify JWT
// @Param headers x-token "JWT" body raw json {name: "博客名字", title: "博客标题", text: "简述", tag: [..."标签"]}
// @Success 200
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

// @Summary 创建博客博客的详细信息
// @Router /system/blog/createinfo [post]
// @Verify JWT
// @Param headers x-token "JWT" body raw json {name: "博客名字", title: "博客标题", text: "简述", tag: [..."标签"]}
// @Success 200
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

// @Summary 删除博客的图片和markdown文档
// @Router /system/blog/uploadfile [post]
// @Verify JWT
// @Param headers x-token "JWT" uri "博客名字"
// @Success 200
func (s *BlogApi) DelBlogInfo(c *gin.Context) {
	param := c.Param("blog")

	err := service.ServiceGroupApp.SystemServiceGroup.BlogServiceEx.DelBlogInfo(param)
	if err != nil {
		response.FailWithMessage("DelBlogInfoApi SQL Error", c)
	}

	err = os.RemoveAll(utils.GetBlogPath(param))
	if err != nil {
		response.FailWithMessage("DelBlogInfoApi file operations Error", c)
	}
	response.Ok(c)
}
