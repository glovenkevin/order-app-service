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

type UserRepo struct {
	*bun.DB
	Log logger.LoggerInterface
}

func NewUserRepo(log logger.LoggerInterface, db *bun.DB) domain.UserRepoInterface {
	return &UserRepo{Log: log, DB: db}
}

func (u *UserRepo) Begin(ctx context.Context) (*bun.Tx, error) {
	select {
	case <-ctx.Done():
		return nil, error_helper.ContextError(ctx)
	default:
	}

	tx, err := u.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *UserRepo) InsertUser(ctx context.Context, tx *bun.Tx, user *entity.User) error {
	select {
	case <-ctx.Done():
		return error_helper.ContextError(ctx)
	default:
	}

	_, err := tx.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	select {
	case <-ctx.Done():
		return nil, error_helper.ContextError(ctx)
	default:
	}

	user := new(entity.User)
	err := r.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
