package system

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

type ListRouter struct {
}

func (b *ListRouter) InitListRouter(Router *gin.RouterGroup) {
	listRouter := Router.Group("system")
	var listApi = api.ApiGroupApp.SystemApiGroup.ListApi
	listRouter.GET("list", listApi.GetList)
}
