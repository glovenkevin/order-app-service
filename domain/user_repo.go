package domain

import "order-app/domain/entity"

type UserRepo interface {
	RegisterUser(user *entity.User) error
}
