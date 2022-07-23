package testcase

import (
	"order-app/domain/model"
	error_helper "order-app/pkg/error"
)

type LoginTestCase struct {
	Name   string
	Req    *model.LoginRequest
	Res    *model.LoginResponse
	ErrMsg string
}

func (t *LoginTestCase) DefaultLogin() *LoginTestCase {
	tc := &LoginTestCase{
		Name: "DefaultLogin",
	}

	tc.Req = &model.LoginRequest{
		Email:    "admin@admin.com",
		Password: "admin",
	}
	tc.Res = &model.LoginResponse{
		Message: "User logged in successfully",
	}
	tc.ErrMsg = ""

	return tc
}

func (t *LoginTestCase) PasswordIncorrect() *LoginTestCase {
	tc := &LoginTestCase{
		Name: "PasswordIncorrect",
	}

	tc.Req = &model.LoginRequest{
		Email:    "admin@admin.com",
		Password: "adminf",
	}
	tc.Res = nil
	tc.ErrMsg = error_helper.ErrPasswordIncorrect.Error()

	return tc
}

func (t *LoginTestCase) UserNotFound() *LoginTestCase {
	tc := &LoginTestCase{
		Name: "UserNotFound",
	}

	tc.Req = &model.LoginRequest{
		Email:    "asdf@asdf.com",
		Password: "admin",
	}
	tc.Res = nil
	tc.ErrMsg = error_helper.ErrUserNotFound.Error()

	return tc
}
