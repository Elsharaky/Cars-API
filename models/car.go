package models

import "github.com/lib/pq"

type Car struct {
	ID     uint          `json:"id"`
	Name   string        `json:"name" validate:"required,min=3,max=64"`
	Make   string        `json:"make" validate:"required,min=3,max=64"`
	Model  uint          `json:"model" validate:"required,min=1800,numeric"`
	Color  string        `json:"color" validate:"required,min=3,max=64"`
	Speed  pq.Int32Array `json:"speed" validate:"required,len=2,dive,min=0,max=1000"`
	TypeID uint          `json:"type_id,omitempty" validate:"required,numeric"`
	Type   string        `json:"type,omitempty"`
}
