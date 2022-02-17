package system

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

type BlogRouter struct {
}

func (b *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	listRouter := Router.Group("system")
	var listApi = api.ApiGroupApp.SystemApiGroup.ListApi
	{
		listRouter.GET("blog/:blog", listApi.GetBlog)
	}
}
