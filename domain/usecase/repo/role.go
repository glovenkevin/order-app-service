package repo

import (
	"context"
	"database/sql"
	"order-app/domain"
	"order-app/domain/entity"
	error_helper "order-app/pkg/error"
	"order-app/pkg/logger"

	"github.com/uptrace/bun"
)

type RoleRepo struct {
	*bun.DB
	Log logger.LoggerInterface
}

func NewRoleRepo(log logger.LoggerInterface, db *bun.DB) domain.RoleRepoInterface {
	return &RoleRepo{DB: db, Log: log}
}

func (r *RoleRepo) GetRoleByCode(ctx context.Context, code string) (*entity.Role, error) {
	tracert := "repo.RoleRepo.GetRoleByCode"
	select {
	case <-ctx.Done():
		return nil, error_helper.ContextError(ctx)
	default:
	}

	role := &entity.Role{}
	err := r.DB.NewSelect().Model(role).Where("code = ?", code).Scan(ctx)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		r.Log.Errorf(tracert+" - r.DB.Model(role).Where(code = ?, code).Select: %w", err)
		return nil, err
	}

	return role, nil
}
