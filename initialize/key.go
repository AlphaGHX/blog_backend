package initialize

import (
	"blog/global"
	"fmt"
	"io/ioutil"
	"os"
)

func InitKey() []byte {
	file, err := os.Open(global.CONFIG.Jwt.SecretPath)
	if err != nil {
		panic(fmt.Errorf("ERROR\tjwtKey打开时错误: %v", err))
	}
	defer file.Close()
	sha256key, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("ERROR\tjwtKey读取时错误: %v", err))
	}
	return sha256key
}
