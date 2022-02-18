package initialize

// import (
// 	"blog/global"
// 	"blog/models"
// 	"fmt"
// 	"io/ioutil"

// 	"github.com/spf13/viper"
// )

// func ReadBlog() map[string]models.List {
// 	var BlogList = make(map[string]models.List)
// 	v := viper.New()
// 	files, err := ioutil.ReadDir(global.BLOG_HOME)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, f := range files {
// 		v.SetConfigFile(global.BLOG_HOME + f.Name() + global.BLOG_CONFIG)
// 		v.SetConfigType("toml")
// 		if err := v.ReadInConfig(); err != nil {
// 			panic(fmt.Errorf("fatal error in reading config file: %v", err))
// 		}
// 		BlogItem := models.List{
// 			Id:     v.GetString("id"),
// 			Title:  v.GetString("title"),
// 			Text:   v.GetString("text"),
// 			TopImg: "/system/list/img/" + v.GetString("id"),
// 		}
// 		BlogList[v.GetString("id")] = BlogItem
// 	}
// 	return BlogList
// }
