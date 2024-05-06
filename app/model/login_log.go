package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"github.com/google/uuid"
	"time"
)

type LoginType string

const (
	Login  LoginType = "login"
	Logout LoginType = "logout"
)

type LoginLog struct {
	ID        string    `gorm:"primaryKey"` // UUID
	UserID    string    `gorm:"not null" json:"user_id"`
	LoginType LoginType `gorm:"not null" json:"login_type"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func (l *LoginLog) CreateLoginLog() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	l.ID = uuid.New().String()

	return db.Create(l).Error
}
