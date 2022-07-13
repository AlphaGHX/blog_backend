package initialize

import (
	"blog/global"
	"os"
	fp "path/filepath"
)

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
