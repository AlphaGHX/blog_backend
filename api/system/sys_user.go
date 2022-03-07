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

func (s *UserApi) VerifyToken(c *gin.Context) {
	response.OkWithMessage("验证通过", c)
}
