package usecase

import (
	"context"
	"order-app/domain"
	"order-app/domain/entity"
	"order-app/domain/model"
	error_helper "order-app/pkg/error"
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

func NewAuthUseCase(log logger.LoggerInterface, userRepo domain.UserRepoInterface, roleRepo domain.RoleRepoInterface) Auther {
	return &AuthUseCase{Log: log, UserRepo: userRepo, RoleRepo: roleRepo}
}

func (u *AuthUseCase) Register(ctx context.Context, user *entity.User) error {
	tracestr := "usecase.AuthUseCase.Register"
	select {
	case <-ctx.Done():
		return error_helper.ContextError(ctx)
	default:
	}

	err := u.ValidateNewUser(ctx, &model.RegisterRequest{Email: user.Email})
	if err != nil {
		u.Log.Errorf(tracestr+" - u.ValidateNewUser: %w", err)
		return err
	}

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

	err = u.UserRepo.InsertUser(ctx, tx, user)
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

func (u *AuthUseCase) ValidateNewUser(ctx context.Context, req *model.RegisterRequest) error {
	tracestr := "usecase.AuthUseCase.ValidateNewUser"
	select {
	case <-ctx.Done():
		return error_helper.ContextError(ctx)
	default:
	}

	user, err := u.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		u.Log.Errorf(tracestr+" - u.UserRepo.GetUserByEmail: %w", err)
		return err
	}

	if user.Email != "" {
		return error_helper.ErrEmailAlreadyExists
	}

	return nil
}

func (u *AuthUseCase) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	tracestr := "usecase.AuthUseCase.Login"

	usr, err := u.UserRepo.GetUserByEmail(ctx, req.Email)
	if usr == nil {
		u.Log.Warnf(tracestr+" - u.UserRepo.GetUserByEmail: User with email %w is not exist", req.Email)
		return nil, error_helper.ErrUserNotFound
	}
	if err != nil {
		u.Log.Errorf(tracestr+" - u.UserRepo.GetUserByEmail: %w", err)
		return nil, err
	}

	if usr.Email == "" {
		return nil, error_helper.ErrUserNotFound
	}

	pwb := []byte(req.Password)
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), pwb)
	if err != nil {
		u.Log.Errorf(tracestr+" - bcrypt.CompareHashAndPassword: %w", err)
		return nil, error_helper.ErrPasswordIncorrect
	}

	return &model.LoginResponse{
		Token:   "",
		Message: "User logged in successfully",
	}, nil
}
