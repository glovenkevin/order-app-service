package repo

import (
	"context"
	"order-app/domain/entity"
	"order-app/pkg/logger"

	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	*pg.DB
	Log logger.LoggerInterface
}

func NewUserRepo(log logger.LoggerInterface, db *pg.DB) *UserRepo {
	return &UserRepo{Log: log, DB: db}
}

func (u *UserRepo) Begin(ctx context.Context) (*pg.Tx, error) {
	return u.DB.BeginContext(ctx)
}

func (r *UserRepo) RegisterUser(tx *pg.Tx, user *entity.User) error {
	_, err := tx.Model(user).Insert()
	return err
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.ModelContext(ctx, &user).Where("email = ?", email).Select()
	return &user, err
}
