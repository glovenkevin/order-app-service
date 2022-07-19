package testcase

import "order-app/domain/entity"

type RegisterUserTestCase struct {
	Name string
	Req  *entity.User
	Res  error
}

func (t *RegisterUserTestCase) DefaultParam() *RegisterUserTestCase {
	tc := &RegisterUserTestCase{
		Name: "DefaultParam",
	}

	tc.Req = &entity.User{
		Name:        "Kevin",
		Email:       "kevin@kevin.com",
		Password:    "test123",
		FcmToken:    "asdfasdf",
		PhoneNumber: "123123",
		IsBlocked:   false,
	}

	tc.Res = nil

	return tc
}
