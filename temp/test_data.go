package temp

import (
	"blog/global"
	"blog/models"
)

func InsertTestData() {
	data := []models.Blog{
		{
			Name:  "about-alpha-blog",
			Title: "关于AlphaBlog",
			Text:  "至少又想起来了，不是吗？",
		},
		{
			Name:  "fist",
			Title: "第一条博客",
			Text:  "正在思考写啥",
		},
		{
			Name:  "test",
			Title: "Markdown测试",
			Text:  "Markdown测试,Markdown测试",
		},
	}

	global.GROM.Create(&data)
}
