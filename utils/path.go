package utils

import (
	"blog/global"
	fp "path/filepath"
)

var local = &global.CONFIG.Local

// 博客文件夹路径
func GetBlogPath(blogName string) string {
	return fp.Join(local.BloghomeDir, blogName)
}

// 博客markdown文件路径
func GetBlogMarkdownPath(blogName string) string {
	return fp.Join(local.BloghomeDir, blogName, local.MarkdownFile)
}

// 博客顶部图片文件路径
func GetBlogTopimgPath(blogName string) string {
	return fp.Join(local.BloghomeDir, blogName, local.TopimgFile)
}
