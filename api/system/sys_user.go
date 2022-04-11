package system

import (
	"blog/models"
	"blog/models/request"
	"blog/models/response"
	"blog/service"
	"blog/utils"

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
