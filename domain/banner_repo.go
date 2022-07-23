package domain

import (
	"context"
	"order-app/domain/entity"
)

type BannerRepoInterface interface {
	Find(ctx context.Context) ([]*entity.Banner, error)
}
