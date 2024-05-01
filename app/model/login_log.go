package model

import (
	"github.com/google/uuid"
	"time"
)

type LoginType string

const (
	Login  LoginType = "login"
	Logout LoginType = "logout"
)

type LoginLog struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    string    `gorm:"not null" json:"user_id"`
	LoginType LoginType `gorm:"not null" json:"login_type"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (l LoginLog) CreateLoginLog() {

}
