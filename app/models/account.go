package models

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID           uuid.UUID     `gorm:"primaryKey" json:"id"` // UUID
	UserID       string        `gorm:"not null" json:"user_id"`
	Balance      int           `gorm:"not null;default:0" json:"balance"` // 잔고 (음수 금액 불가능, 소숫점 사용 불가능)
	CreatedAt    time.Time     `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	Transactions []Transaction `gorm:"foreignKey:AccountID" json:"transactions,omitempty"`
}
