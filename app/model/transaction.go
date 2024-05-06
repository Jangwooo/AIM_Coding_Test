package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"github.com/google/uuid"
	"time"
)

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
)

type Transaction struct {
	ID           string          `gorm:"primaryKey" json:"id"` // UUID
	AccountID    string          `gorm:"not null" json:"account_id"`
	Amount       uint            `gorm:"not null" json:"amount"`
	Type         TransactionType `gorm:"not null" json:"type"` // 입금, 출금
	AfterBalance uint            `gorm:"not null" json:"after_balance"`
	CreatedAt    time.Time       `gorm:"not null" json:"created_at"`
}

func (t *Transaction) CreateTransaction() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	t.ID = uuid.New().String()
	return db.Create(t).Error
}
