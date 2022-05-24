package initialize

import (
	"blog/middleware"
	"blog/routers"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()
	Router.SetTrustedProxies(nil)

	systemRouter := routers.RouterGroupApp.System

	PublicGroup := Router.Group("")
	{
		systemRouter.InitListRouter(PublicGroup)
		systemRouter.InitBlogRouter(PublicGroup)
		systemRouter.InitTagRouter(PublicGroup)
		systemRouter.InitUserRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JwtAuth())
	{
		systemRouter.InitBlogRouterEx(PrivateGroup)
		systemRouter.InitUserRouterEx(PrivateGroup)
	}

	return Router
}
