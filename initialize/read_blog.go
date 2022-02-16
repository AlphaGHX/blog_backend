package initialize

import (
	"blog/global"
	"blog/models"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func ReadBlog() map[string]models.List {
	var BlogList = make(map[string]models.List)
	v := viper.New()
	global.BLOG_LIST = make(map[string]models.List)
	files, err := ioutil.ReadDir("./blog")
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		v.SetConfigFile("./blog/" + f.Name() + "/config.toml")
		v.SetConfigType("toml")
		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal error in reading config file: %v", err))
		}
		BlogItem := models.List{
			Id:    v.GetString("id"),
			Title: v.GetString("title"),
			Text:  v.GetString("text"),
		}
		BlogList[v.GetString("id")] = BlogItem
	}
	return BlogList
}
