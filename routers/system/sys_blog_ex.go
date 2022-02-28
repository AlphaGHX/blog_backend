package system

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

type BlogRouterEx struct {
}

func (b *BlogRouter) InitBlogRouterEx(Router *gin.RouterGroup) {
	blogRouter := Router.Group("system")
	var blogApi = api.ApiGroupApp.SystemApiGroup.BlogApi
	{
		blogRouter.POST("blog/uploadfile", blogApi.PostBlogFile)
		blogRouter.POST("blog/uploadinfo", blogApi.PostBlogInfo)
		blogRouter.POST("blog/createinfo", blogApi.CreateBlogInfo)
		blogRouter.POST("blog/del/:blog", blogApi.DelBlogInfo)
	}
}
