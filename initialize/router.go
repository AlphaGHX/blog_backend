package initialize

import (
	"blog/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(cors.Default())

	systemRouter := routers.RouterGroupApp.System

	PublicGroup := Router.Group("")
	{
		systemRouter.InitListRouter(PublicGroup)
	}

	return Router
}
