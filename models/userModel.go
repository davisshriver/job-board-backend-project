package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	FirstName *string   `json:"first_name" validate:"required,min=2,max=100"`
	LastName  *string   `json:"last_name" validate:"required,min=2,max=100"`
	Password  *string   `json:"password" validate:"required,min=6,max=100"`
	Email     *string   `json:"email" validate:"email,required"`
	Phone     *string   `json:"phone" validate:"required,min=1,max=10"`
	UserType  *string   `json:"user_type" validate:"required,eq=ADMIN|eq=USER"` // Can be an Admin or User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UID       string    `json:"user_id"`
}
