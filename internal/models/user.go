package models

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" validate:"required,min=5,max=100"`
	Password  string    `json:"-" validate:"required,min=20,max=40"`
	Role      string    `json:"role" validate:"oneof=USER ADMIN"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
