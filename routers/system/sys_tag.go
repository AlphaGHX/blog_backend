package system

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

func (b *TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	tagRouter := Router.Group("system")
	var tagApi = api.ApiGroupApp.SystemApiGroup.TagApi
	{
		tagRouter.GET("tag/:tag", tagApi.GetBlogUseTag)
		tagRouter.GET("tag", tagApi.GetTagList)
	}
}
