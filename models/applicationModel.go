package models

import "time"

type Application struct {
	ApplicationID int       `gorm:"primaryKey;autoIncrement"`
	PostID        int       `json:"PostID" validate:"required"`
	UserID        int       `json:"description" validate:"required,min=2,max=300"`
	Experience    string    `json:"requirements" validate:"required,min=2,max=300"`
	CreatedBy     string    `json:"created_by" validate:"required"`
	Address       string    `json:"location" validate:"required"`
	Wage          int       `json:"wage" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"expires_at"`
}
