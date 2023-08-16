package inputs

import "time"

type UserUpdate struct {
	FirstName *string    `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName  *string    `json:"last_name" validate:"omitempty,min=2,max=100"`
	Password  *string    `json:"password" validate:"omitempty,min=6,max=100"`
	Email     *string    `json:"email" validate:"omitempty,email"`
	Phone     *string    `json:"phone" validate:"omitempty,min=1,max=10"`
	UserType  *string    `json:"user_type" validate:"omitempty,eq=ADMIN|eq=USER"`
	UpdatedAt *time.Time `json:"updated_at" validate:"omitempty"`
}
