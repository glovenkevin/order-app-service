package domain

import (
	"context"
	"order-app/domain/entity"

	"github.com/go-pg/pg/v10"
)

type UserRepoInterface interface {
	Begin(ctx context.Context) (*pg.Tx, error)
	InsertUser(tx *pg.Tx, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
