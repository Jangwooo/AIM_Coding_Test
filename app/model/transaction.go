package model

import (
	"github.com/google/uuid"
	"time"
)

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
)

type Transaction struct {
	ID        uuid.UUID       `gorm:"primaryKey" json:"id"` // UUID
	AccountID string          `gorm:"not null" json:"account_id"`
	Amount    int64           `gorm:"not null" json:"amount"`
	Type      TransactionType `gorm:"not null" json:"type"` // 입금, 출금
	CreatedAt time.Time       `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (t Transaction) CreateTransaction() {

}
