## blog_backend
Golang实现的一个简单的博客后端

## 使用到的技术
* Gin
* GORM
* SQLite
* JWT

## 已有功能
* 博客文章增删改查
* 根据tag查找对应文章
* 访问量统计
* 文件上传下载
* 登录
* 使用jwt免密码登录

## 准备补充的功能
* 分页

## 前端
[alpha-blog-next](https://github.com/AlphaGHX/alpha-blog-next)

## 运行
1. 重命名config.sample.toml为config.toml
2. 根据config.toml中的提示生成公钥私钥（这用于免密码登录）
3. 配置config.toml中的管理员用户名密码
4. go mod tidy
5. go run       (开发环境)
6. go build     (生产环境)
