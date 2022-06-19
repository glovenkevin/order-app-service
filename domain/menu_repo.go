package domain

import (
	"context"
	"order-app/domain/entity"
	"order-app/domain/model"
)

type MenuRepoInterface interface {
	Find(ctx context.Context, req *model.MenuRequest) ([]*entity.Menu, error)
}
