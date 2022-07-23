package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"table:menues"`

	ID          uuid.UUID `bun:"id"`
	Name        string    `bun:"name"`
	Stock       int32     `bun:"stock"`
	Description *string   `bun:"description"`
	Price       float64   `bun:"price"`
	ImageUrl    string    `bun:"image_url"`
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
}
