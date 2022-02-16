package response

import "blog/models"

type ListResponse struct {
	List map[string]models.List `json:"list"`
}
