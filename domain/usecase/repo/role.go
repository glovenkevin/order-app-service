package repo

import (
	"context"
	"order-app/domain"
	"order-app/domain/entity"
	error_helper "order-app/pkg/error"
	"order-app/pkg/logger"

	"github.com/go-pg/pg/v10"
)

type RoleRepo struct {
	*pg.DB
	Log logger.LoggerInterface
}

func NewRoleRepo(log logger.LoggerInterface, db *pg.DB) domain.RoleRepoInterface {
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
	err := r.DB.ModelContext(ctx, role).Where("code = ?", code).Select()
	if err != nil {
		r.Log.Errorf(tracert+" - r.DB.Model(role).Where(code = ?, code).Select: %w", err)
		return nil, err
	}

	return role, nil
}
