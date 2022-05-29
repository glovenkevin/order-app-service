package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `pg:"id"`
	RoleID      uuid.UUID `pg:"role_id"`
	Name        string    `pg:"name"`
	Email       string    `pg:"email"`
	Password    string    `pg:"password"`
	PhoneNumber string    `pg:"phone_number"`
	FcmToken    string    `pg:"fcm_token"`
	IsBlocked   bool      `pg:"is_blocked"`
	CreatedAt   time.Time `pg:"created_at"`
	UpdatedAt   time.Time `pg:"updated_at"`
}
