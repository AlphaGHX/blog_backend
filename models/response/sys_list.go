package response

import "blog/models"

type ListResponse struct {
	List []models.List `json:"list"`
}
