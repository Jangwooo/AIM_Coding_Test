package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID           string        `gorm:"primaryKey" json:"id"`
	UserID       string        `gorm:"not null" json:"user_id"`
	Balance      uint          `gorm:"not null;default:0" json:"balance"` // 잔고 (음수 금액 불가능, 소숫점 사용 불가능)
	CreatedAt    time.Time     `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	Transactions []Transaction `gorm:"foreignKey:AccountID" json:"transactions,omitempty"`
}

func (a *Account) CreateAccount() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Create(a).Error
}

func (a *Account) GetAccountByID() (*Account, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	err = db.First(a).Error
	return a, err
}

func (a *Account) GetAccountList() ([]Account, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var accounts []Account

	err = db.Find(&accounts, "user_id = ?", a.UserID).Order("created_at desc").Error

	return accounts, err
}

func (a *Account) Deposit(amount uint) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Model(a).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func (a *Account) Withdraw(amount uint) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Model(a).Update("balance", gorm.Expr("balance - ?", amount)).Error
}
