package system

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

type UserRouterEx struct {
}

func (b *UserRouterEx) InitUserRouterEx(Router *gin.RouterGroup) {
	userRouter := Router.Group("system")
	var userApi = api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("token-verify", userApi.VerifyToken)
	}
}
