package testcase

import (
	"context"
	"order-app/domain/usecase"
	"order-app/domain/usecase/repo"
	"order-app/pkg/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (t *LoginTestCase) GetScenarios() []*LoginTestCase {
	return []*LoginTestCase{
		t.DefaultLogin(),
		t.PasswordIncorrect(),
		t.UserNotFound(),
	}
}

func TestLogin(t *testing.T) {
	tests := new(LoginTestCase).GetScenarios()

	tc := test.NewTestCase(t)
	defer tc.TearDown()

	ctx := context.Background()

	userRepo := repo.NewUserRepo(tc.Log, tc.Db)
	roleRepo := repo.NewRoleRepo(tc.Log, tc.Db)
	uc := usecase.NewAuthUseCase(tc.Log, userRepo, roleRepo)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := uc.Login(ctx, test.Req)

			if test.ErrMsg != "" {
				assert.Nil(t, res)
				assert.Contains(t, err.Error(), test.ErrMsg)
				return
			}

			assert.NoError(t, err)
			if res != nil {
				assert.Equal(t, test.Res.Message, res.Message)
			}
		})
	}
}
