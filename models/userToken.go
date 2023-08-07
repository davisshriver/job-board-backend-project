package models

import (
	"time"
)

type UserToken struct {
	TokenID      uint      `gorm:"primaryKey"`
	UserId       string    `gorm:"index"`
	Token        string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
