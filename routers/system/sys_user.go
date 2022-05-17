package system

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (b *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("system")
	var userApi = api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.GET("get-admin-info", userApi.GetAdminInfo)
		userRouter.GET("avatar", userApi.GetAdminAvatar)
		userRouter.POST("user-verify", userApi.VerifyUser)
	}
}
