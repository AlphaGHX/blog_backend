package initialize

import (
	"blog/global"
	"os"
	fp "path/filepath"
)

/*
./data: 数据目录
./data/blogs: 博客目录
./data/blog-name/markdown: 博客的 markdown 文件
./data/blog-name/topimg: 博客的顶部图片

./data/user: 管理员目录
./data/user/admin-avatar: 管理员头像
*/

func InitFile() {
	Local := &global.CONFIG.Local
	Local.DataDir = "data"
	Local.BloghomeDir = fp.Join(Local.DataDir, "blogs")
	Local.UserhomeDir = fp.Join(Local.DataDir, "user")
	Local.AdminAvatarFile = "admin-avatar"
	Local.MarkdownFile = "markdown"
	Local.TopimgFile = "topimg"

	os.Mkdir(Local.DataDir, os.ModePerm)
	os.Mkdir(Local.BloghomeDir, os.ModePerm)
	os.Mkdir(Local.UserhomeDir, os.ModePerm)
}
