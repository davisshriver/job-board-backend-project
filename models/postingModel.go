package models

import (
	"time"
)

type JobPost struct {
	PostID       int       `gorm:"primaryKey;autoIncrement"`
	Role         string    `json:"role" validate:"required,min=2,max=100"`
	Description  string    `json:"description" validate:"required,min=2,max=300"`
	Requirements string    `json:"requirements" validate:"required,min=2,max=300"`
	CreatedBy    string    `json:"created_by" validate:"required"`
	Location     string    `json:"location" validate:"required"`
	Wage         int       `json:"wage" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type JobPostUpdate struct {
	Role         *string    `json:"role" validate:"omitempty,min=2,max=100"`
	Description  *string    `json:"description" validate:"omitempty,min=2,max=300"`
	Requirements *string    `json:"requirements" validate:"omitempty,min=2,max=300"`
	Wage         *int       `json:"wage" validate:"omitempty"`
	Expires_At   *time.Time `json:"expires_at" validate:"omitempty"`
}
