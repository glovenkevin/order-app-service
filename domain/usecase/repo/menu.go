package repo

import (
	"context"
	"order-app/domain"
	"order-app/domain/entity"
	"order-app/domain/model"
	error_helper "order-app/pkg/error"
	"order-app/pkg/logger"

	"github.com/uptrace/bun"
)

type MenuRepo struct {
	*bun.DB
	Log logger.LoggerInterface
}

func NewMenuRepo(db *bun.DB, log logger.LoggerInterface) domain.MenuRepoInterface {
	return &MenuRepo{DB: db, Log: log}
}

func (r *MenuRepo) Find(ctx context.Context, req *model.MenuRequest) ([]*entity.Menu, error) {
	select {
	case <-ctx.Done():
		return nil, error_helper.ContextError(ctx)
	default:
	}

	var menus []*entity.Menu
	q := r.DB.NewSelect().Model(&menus)

	if req.Name != "" {
		q = q.Where("name = ?", req.Name)
	}
	if req.Description != "" {
		q = q.Where("description like %?%", req.Description)
	}
	if req.Price != 0 {
		q = q.Where("price = ?", req.Price)
	}

	if req.Limit != 0 && req.Offset != 0 {
		q = q.Limit(int(req.Limit)).Offset(int(req.Offset))
	}

	err := q.Scan(ctx)
	if err != nil {
		r.Log.Errorf("Find menu failed: %v", err)
		return nil, err
	}

	return menus, nil
}
