package usecase

import (
	"context"
	"order-app/domain"
	"order-app/domain/entity"
	"order-app/domain/model"
	"order-app/pkg/logger"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	Log      logger.LoggerInterface
	UserRepo domain.UserRepoInterface
	RoleRepo domain.RoleRepoInterface
}

func NewAuthUseCase(log logger.LoggerInterface, userRepo domain.UserRepoInterface, roleRepo domain.RoleRepoInterface) *AuthUseCase {
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

	tx, err := u.UserRepo.Begin(ctx)
	if err != nil {
		u.Log.Errorf(tracestr+" - u.UserRepo.Begin: %w", err)
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

	err = u.UserRepo.RegisterUser(tx, user)
	if err != nil {
		tx.Rollback()
		u.Log.Errorf(tracestr+" - u.UserRepo.RegisterUser: %w", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		u.Log.Errorf(tracestr+" - tx.Commit: %w", err)
		return err
	}

	return nil
}

func (u *AuthUseCase) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	tracestr := "usecase.AuthUseCase.Login"

	user, err := u.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		u.Log.Errorf(tracestr+" - u.UserRepo.GetUserByEmail: %w", err)
		return nil, err
	}

	pwb := []byte(req.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), pwb)
	if err != nil {
		u.Log.Errorf(tracestr+" - bcrypt.CompareHashAndPassword: %w", err)
		return nil, err
	}

	return &model.LoginResponse{
		Token:   "",
		Message: "User logged in successfully",
	}, nil
}
