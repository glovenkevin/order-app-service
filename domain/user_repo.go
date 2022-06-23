package domain

import (
	"context"
	"order-app/domain/entity"

	"github.com/uptrace/bun"
)

type UserRepoInterface interface {
	Begin(ctx context.Context) (*bun.Tx, error)
	InsertUser(ctx context.Context, tx *bun.Tx, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
