package request

import "blog/models/response"

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SetAdminInfo struct {
	Avatar string `json:"avatar" binding:"required"`
	response.GetAdminInfo
}
