package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID          uuid.UUID `bun:"id"`
	RoleID      uuid.UUID `bun:"role_id"`
	Name        string    `bun:"name"`
	Email       string    `bun:"email"`
	Password    string    `bun:"password"`
	PhoneNumber string    `bun:"phone_number"`
	FcmToken    string    `bun:"fcm_token"`
	IsBlocked   bool      `bun:"is_blocked"`
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
}
