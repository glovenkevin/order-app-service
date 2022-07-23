package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Banner struct {
	bun.BaseModel `bun:"table:banners"`

	ID          uuid.UUID  `bun:"id"`
	Name        string     `bun:"name"`
	Description *string    `bun:"description"`
	ImageUrl    string     `bun:"image_url"`
	IsDeleted   bool       `bun:"is_deleted"`
	IsShow      bool       `bun:"is_show"`
	Seq         int32      `bun:"seq"`
	CreatedAt   time.Time  `bun:"created_at"`
	UpdatedAt   *time.Time `bun:"updated_at"`
}
