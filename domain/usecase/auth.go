package usecase

import (
	"context"
	"order-app/domain"
	"order-app/domain/entity"
	"order-app/pkg/logger"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	Log      logger.ILogger
	UserRepo domain.UserRepo
	RoleRepo domain.RoleRepo
}

func NewAuthUseCase(log logger.ILogger, userRepo domain.UserRepo, roleRepo domain.RoleRepo) *AuthUseCase {
	return &AuthUseCase{Log: log, UserRepo: userRepo, RoleRepo: roleRepo}
}

func (u *AuthUseCase) Register(ctx context.Context, user *entity.User) error {
	tracestr := "usecase.AuthUseCase.Register"

	pwb := []byte(user.Password)
	hpw, err := bcrypt.GenerateFromPassword(pwb, bcrypt.DefaultCost)
	if err != nil {
		u.Log.Errorf(tracestr+" - bcrypt.GenerateFromPassword: %w", err)
		return err
	}

	role, err := u.RoleRepo.GetRoleByCode(ctx, "USR")
	if err != nil {
		u.Log.Errorf(tracestr+" - u.RoleRepo.GetRoleByCode: %w", err)
		return err
	}

	user.ID = uuid.New()
	user.RoleID = role.ID
	user.Password = string(hpw)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return u.UserRepo.RegisterUser(user)
}

// func (u *AuthUseCase) Login(user *) error {
// 	return u.UserRepo.RegisterUser(user)
// }
