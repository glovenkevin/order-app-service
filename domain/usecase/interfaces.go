package usecase

import (
	"context"
	"order-app/domain/entity"
	"order-app/domain/model"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Auth -.
	Auther interface {
		ValidateNewUser(ctx context.Context, req *model.RegisterRequest) error
		Register(ctx context.Context, user *entity.User) error
		Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	}

	// Menu -.
	Menuer interface {
		GetMenues(ctx context.Context, req *model.MenuRequest) ([]*model.MenuResponse, error)
	}

	// Banners -.

	Banner interface {
		GetBanners(ctx context.Context) ([]*model.BannerResponse, error)
	}
)
