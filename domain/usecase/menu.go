package usecase

import (
	"context"
	"order-app/domain"
	"order-app/domain/model"
	error_helper "order-app/pkg/error"
	"order-app/pkg/logger"
	"order-app/pkg/variable"
)

type MenuUseCase struct {
	Log      logger.LoggerInterface
	MenuRepo domain.MenuRepoInterface
}

func NewMenuUsecase(log logger.LoggerInterface, menuRepo domain.MenuRepoInterface) Menuer {
	return &MenuUseCase{Log: log, MenuRepo: menuRepo}
}

func (u *MenuUseCase) GetMenues(ctx context.Context, req *model.MenuRequest) ([]*model.MenuResponse, error) {
	tracestr := "usecase.AuthUseCase.Register"
	select {
	case <-ctx.Done():
		return nil, error_helper.ContextError(ctx)
	default:
	}

	menus, err := u.MenuRepo.Find(ctx, req)
	if err != nil {
		u.Log.Errorf("%v - error: %v", tracestr, err)
		return nil, err
	}

	ret := make([]*model.MenuResponse, len(menus))
	for i, m := range menus {
		ret[i] = &model.MenuResponse{
			ID:          m.ID.String(),
			Name:        m.Name,
			Price:       m.Price,
			Description: variable.GetString(m.Description),
			ImageUrl:    m.ImageUrl,
			Stock:       m.Stock,
		}
	}

	return ret, nil
}
