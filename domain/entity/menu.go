package entity

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID          uuid.UUID `pg:"id"`
	Name        string    `pg:"name"`
	Stock       int32     `pg:"stock"`
	Description string    `pg:"description"`
	Price       float64   `pg:"price"`
	ImageUrl    string    `pg:"image_url"`
	CreatedAt   time.Time `pg:"created_at"`
	UpdatedAt   time.Time `pg:"updated_at"`
}
