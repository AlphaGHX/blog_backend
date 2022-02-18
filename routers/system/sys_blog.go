package system

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

type BlogRouter struct {
}

func (b *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	blogRouter := Router.Group("system")
	var blogApi = api.ApiGroupApp.SystemApiGroup.ListApi
	{
		blogRouter.GET("blog/:blog", blogApi.GetBlog)
		blogRouter.GET("blog-info/:blog", blogApi.GetBlogInfo)
	}
}
