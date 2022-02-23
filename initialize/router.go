package initialize

import (
	"blog/middleware"
	"blog/routers"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	fmt.Println("INFO\t启动路由...")
	Router := gin.Default()
	Router.Use(cors.Default())

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

	fmt.Println("DONE\t路由启动成功")
	return Router
}
