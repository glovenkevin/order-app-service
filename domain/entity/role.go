package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID `json:"id" pg:"id"`
	Code      string    `json:"code" pg:"code"`
	Name      string    `json:"name" pg:"name"`
	CreatedAt time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
}
