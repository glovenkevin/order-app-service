package usecase

import (
	"context"
	"order-app/domain"
	"order-app/domain/model"
	error_helper "order-app/pkg/error"
	"order-app/pkg/logger"
)

type BannerUseCase struct {
	Log        logger.LoggerInterface
	BannerRepo domain.BannerRepoInterface
}

func NewBannerUsecase(log logger.LoggerInterface, repo domain.BannerRepoInterface) Banner {
	return &BannerUseCase{Log: log, BannerRepo: repo}
}

func (b *BannerUseCase) GetBanners(ctx context.Context) ([]*model.BannerResponse, error) {
	select {
	case <-ctx.Done():
		return nil, error_helper.ContextError(ctx)
	default:
	}

	mm, err := b.BannerRepo.Find(ctx)
	if err != nil {
		return nil, err
	}

	ret := make([]*model.BannerResponse, len(mm))
	for i, m := range mm {
		ret[i] = &model.BannerResponse{
			Name:        m.Name,
			Description: m.Description,
			ImageUrl:    m.ImageUrl,
			Seq:         m.Seq,
			IsShow:      m.IsShow,
		}
	}

	return ret, nil
}
