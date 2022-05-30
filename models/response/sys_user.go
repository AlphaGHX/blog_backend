package response

import "blog/models"

type GetAdminInfo struct {
	Nickname string          `json:"nickname"`
	Describe string          `json:"describe"`
	MyLinks  []models.MyLink `json:"my-links"`
}
