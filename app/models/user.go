package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID    `gorm:"primaryKey" json:"id"`
	Username   string       `gorm:"not null;unique" json:"username"`
	Password   string       `gorm:"not null" json:"password"`
	CreatedAt  time.Time    `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	LoginLogs  *[]LoginLog  `gorm:"foreignKey:UserID" json:"login_logs,omitempty"`
	Portfolios *[]Portfolio `gorm:"foreignKey:UserID" json:"portfolios,omitempty"`
	Account    *Account     `gorm:"foreignKey:UserID" json:"account,omitempty"`
}
