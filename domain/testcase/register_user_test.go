package testcase

import (
	"context"
	"order-app/domain/usecase"
	"order-app/domain/usecase/repo"
	"order-app/pkg/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (t *RegisterUserTestCase) GetScenarios() []*RegisterUserTestCase {
	return []*RegisterUserTestCase{
		t.DefaultParam(),
		t.UserAlreadyExist(),
	}
}

func TestRegisterUser(t *testing.T) {
	tests := new(RegisterUserTestCase).GetScenarios()

	tc := test.NewTestCase(t)
	defer tc.TearDown()

	ctx := context.Background()

	userRepo := repo.NewUserRepo(tc.Log, tc.Db)
	roleRepo := repo.NewRoleRepo(tc.Log, tc.Db)
	uc := usecase.NewAuthUseCase(tc.Log, userRepo, roleRepo)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := uc.Register(ctx, test.Req)

			assert.Equal(t, test.Res, err)
		})
	}
}
