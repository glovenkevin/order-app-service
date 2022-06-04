package domain

import (
	"context"
	"order-app/domain/entity"
)

type RoleRepoInterface interface {
	GetRoleByCode(ctx context.Context, code string) (*entity.Role, error)
}
