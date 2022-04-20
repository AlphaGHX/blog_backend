package initialize

import (
	"blog/middleware"
	"blog/routers"
	"fmt"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	fmt.Println("INIT\t Route")
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()
	Router.SetTrustedProxies(nil)
	// Router.Use(cors.Default())

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

	fmt.Println("DONE\t Route")
	return Router
}
