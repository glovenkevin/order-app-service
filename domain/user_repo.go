package domain

import "order-app/domain/entity"

type UserRepoInterface interface {
	RegisterUser(user *entity.User) error
}
