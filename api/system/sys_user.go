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
		response.FailWithDetailed(err, "VerifyUserApi ERROR", c)
		return
	}
	token, err := utils.GetToken(data.Username)
	if err != nil {
		response.FailWithDetailed(err, "GetToken ERROR", c)
	}
	response.OkWithData(gin.H{"token": token}, c)
}
