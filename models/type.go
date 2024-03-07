package models

import "time"

type Type struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name" validate:"required,min=3,max=64"`
}
