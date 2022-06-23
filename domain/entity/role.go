package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	ID        uuid.UUID `json:"id" bun:"id"`
	Code      string    `json:"code" bun:"code"`
	Name      string    `json:"name" bun:"name"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at"`
}
