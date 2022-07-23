package repo

import (
	"context"
	"order-app/domain/entity"
	error_helper "order-app/pkg/error"
	"order-app/pkg/logger"

	"github.com/uptrace/bun"
)

type BannerRepo struct {
	Db  *bun.DB
	Log logger.LoggerInterface
}

func NewBannerRepo(db *bun.DB, log logger.LoggerInterface) *BannerRepo {
	return &BannerRepo{Db: db, Log: log}
}

func (r *BannerRepo) Find(ctx context.Context) ([]*entity.Banner, error) {
	select {
	case <-ctx.Done():
		return nil, error_helper.ContextError(ctx)
	default:
	}

	var mm []*entity.Banner
	err := r.Db.NewSelect().Model(&mm).Scan(ctx)
	if err != nil {
		r.Log.Errorf("Find banner failed: %v", err)
		return nil, err
	}

	return mm, nil
}
