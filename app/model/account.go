package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID           string        `gorm:"primaryKey" json:"id"`
	UserID       string        `gorm:"not null" json:"user_id"`
	Balance      int64         `gorm:"not null;default:0" json:"balance"` // 잔고 (음수 금액 불가능, 소숫점 사용 불가능)
	CreatedAt    time.Time     `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	Transactions []Transaction `gorm:"foreignKey:AccountID" json:"transactions,omitempty"`
}

func (account *Account) OpenAccount() {
	db := database.OpenDBConnection()

	db.Create(account)
}

func (account *Account) UpdateAccountBalance(amount int64) error {
	db := database.OpenDBConnection()

	return db.Update("balance", gorm.Expr("balance + ?", amount)).Where("id = ?", account.ID).Error
}

func (account *Account) GetBalance(accountID string) (int64, error) {
	db := database.OpenDBConnection()

	err := db.Where("account_id = ?", accountID).Find(account).Error

	return account.Balance, err
}

func (account *Account) Deposit(amount int64) error {
	db := database.OpenDBConnection()

	return db.Update("balance", gorm.Expr("balance + ?", amount)).Where("id = ?", account.ID).Error
}

func (account *Account) Withdraw(amount int64) error {
	db := database.OpenDBConnection()

	return db.Update("balance", gorm.Expr("balance - ?", amount)).Where("id = ?", account.ID).Error
}
