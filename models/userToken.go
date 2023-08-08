package models

import (
	"time"
)

type UserToken struct {
	TokenID      uint      `gorm:"primaryKey"`
	Token        string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UserID       int       `gorm:"not null"`
}
