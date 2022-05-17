package response

import "blog/models"

type AdminInfo struct {
	Nickname string          `json:"nickname"`
	MyLinks  []models.MyLink `json:"my-links"`
}
