package models

import (
	"github.com/google/uuid"
	"time"
)

type LoginLog struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    string    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
