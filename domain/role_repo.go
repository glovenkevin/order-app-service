package domain

import (
	"context"
	"order-app/domain/entity"
)

type RoleRepo interface {
	GetRoleByCode(ctx context.Context, code string) (*entity.Role, error)
}
