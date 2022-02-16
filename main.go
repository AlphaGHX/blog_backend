package main

import (
	"blog/initialize"
	"blog/global"
)

func main() {
	global.BLOG_LIST = initialize.ReadBlog()
	Router := initialize.Routers()
	Router.Run(":12900")
}
