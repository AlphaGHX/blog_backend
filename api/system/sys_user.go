package system

import (
	"blog/models"
	"blog/models/request"
	"blog/models/response"
	"blog/service"
	"blog/utils"
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

// @Summary 通过用户名密码验证用户
// @Router /system/user-verify [post]
// @Verify none
// @Param raw JSON username: "用户名" password: "密码"
// @Success 200 token: "JWT token"
func (s *UserApi) VerifyUser(c *gin.Context) {
	var data request.User
	c.ShouldBindJSON(&data)
	user := models.User{
		Username: data.Username,
		Password: data.Password,
	}

	err := service.ServiceGroupApp.SystemServiceGroup.UserService.VerifyUser(user)
	if err != nil {
		response.FailWithMessage("VerifyUserApi SQL ERROR", c)
		return
	}

	token, err := utils.GetToken(data.Username)
	if err != nil {
		response.FailWithMessage("GetToken ERROR", c)
		return
	}

	response.OkWithData(gin.H{"token": token}, c)
}

// @Summary 通过JWT验证权限
// @Router /system/token-verify [post]
// @Verify JWT
// @Param headers x-token "Token"
// @Success 200
func (s *UserApi) VerifyToken(c *gin.Context) {
	response.OkWithMessage("验证通过", c)
}

// @Summary 获取管理员的头像
// @Router /system/avatar [get]
// @Verify none
// @Param none
// @Success 200
func (s *UserApi) GetAdminAvatar(c *gin.Context) {
	c.File(utils.GetAdminAvatarPath())
}

// @Summary 获取管理员信息
// @Router /system/get-admin-info [get]
// @Verify none
// @Param none
// @Success 200 json 管理员信息
func (s *UserApi) GetAdminInfo(c *gin.Context) {
	data, err := service.ServiceGroupApp.SystemServiceGroup.GetAdminInfo()
	if err != nil {
		response.FailWithMessage("GetAdminInfo ERROR", c)
		return
	}

	response.OkWithData(data, c)
}

// @Summary 修改管理员信息
// @Router /system/set-admin-info [post]
// @Verify JWT
// @Param raw JSON Nickname string, Describe string, Avatar string
// @Success 200
func (s *UserApi) SetAdminInfo(c *gin.Context) {
	var data request.SetAdminInfo
	c.ShouldBindJSON(&data)

	byteImg, err := base64.StdEncoding.DecodeString(data.Avatar)
	if err != nil {
		response.FailWithMessage("Base64 decode ERROR", c)
		return
	}

	err = ioutil.WriteFile(utils.GetAdminAvatarPath(), byteImg, os.ModePerm)
	if err != nil {
		response.FailWithMessage("Avatar WriteFile ERROR", c)
		return
	}

	err = service.ServiceGroupApp.SystemServiceGroup.SetAdminInfo(data)
	if err != nil {
		response.FailWithMessage("SetAdminInfoApi SQL ERROR", c)
		return
	}

	response.Ok(c)
}
