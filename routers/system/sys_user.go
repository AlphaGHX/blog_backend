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
		userRouter.POST("user-verify", userApi.VerifyUser)
	}
}
