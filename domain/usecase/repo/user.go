package repo

import (
	"order-app/domain/entity"
	"order-app/pkg/logger"

	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	*pg.DB
	Log logger.ILogger
}

func NewUserRepo(log logger.ILogger, db *pg.DB) *UserRepo {
	return &UserRepo{Log: log, DB: db}
}

func (r *UserRepo) RegisterUser(user *entity.User) error {
	_, err := r.Model(user).Insert()

	return err
}
