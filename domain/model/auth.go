package model

import "order-app/domain/entity"

type (
	LoginRequest struct {
		Email    string `json:"email" example:"test@test.com" binding:"required,email,max=30"`
		Password string `json:"password" example:"asdf" binding:"required,max=20"`
	}

	LoginResponse struct {
		Token   string `json:"token" example:"asdf123"`
		Message string `json:"message" example:"Success"`
	}

	RegisterRequest struct {
		Name        string `json:"name" example:"test" binding:"required,max=50"`
		Email       string `json:"email" example:"test@test.com" binding:"required,email,max=30"`
		Password    string `json:"password" example:"asdf" binding:"required,max=20"`
		PhoneNumber string `json:"phone_number" example:"08123456789" binding:"required,numeric,max=13"`
		IsBlocked   bool   `json:"is_blocked" example:"false"`
		FcmToken    string `json:"fcm_token" example:"asdf123" binding:"required"`
	}
)

func (r *RegisterRequest) ToEntity() *entity.User {
	return &entity.User{
		Name:        r.Name,
		Email:       r.Email,
		Password:    r.Password,
		PhoneNumber: r.PhoneNumber,
		FcmToken:    r.FcmToken,
		IsBlocked:   r.IsBlocked,
	}
}
