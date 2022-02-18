package response

import "blog/models"

type ListResponse struct {
	List []models.Blog `json:"list"`
}
