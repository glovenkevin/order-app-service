package model

import (
	"order-app/domain/entity"

	"github.com/google/uuid"
)

type (
	MenuRequest struct {
		ID          string  `form:"id"`
		Name        string  `form:"name"`
		Description string  `form:"description"`
		Price       float64 `form:"price" binding:"numeric"`
		Limit       int32   `form:"limit" binding:"numeric"`
		Offset      int32   `form:"offset" binding:"numeric"`
	}

	MenuResponse struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Stock       int32   `json:"stock"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		ImageUrl    string  `json:"image_url"`
	}
)

func (r *MenuRequest) ToEntity() *entity.Menu {
	e := entity.Menu{
		Name:        r.Name,
		Description: &r.Description,
		Price:       r.Price,
	}
	if r.ID != "" {
		id, err := uuid.Parse(r.ID)
		if err == nil {
			e.ID = id
		}
	}

	return &e
}
