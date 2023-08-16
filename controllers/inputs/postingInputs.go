package inputs

import "time"

type JobPostUpdate struct {
	Role         *string    `json:"role" validate:"omitempty,min=2,max=100"`
	Description  *string    `json:"description" validate:"omitempty,min=2,max=300"`
	Requirements *string    `json:"requirements" validate:"omitempty,min=2,max=300"`
	Wage         *int       `json:"wage" validate:"omitempty"`
	Expires_At   *time.Time `json:"expires_at" validate:"omitempty"`
}
