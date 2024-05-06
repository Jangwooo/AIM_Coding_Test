package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"time"
)

type StockTransaction struct {
	ID        string    `gorm:"primary_key" json:"id"`
	UserID    string    `gorm:"not null" json:"user_id"`
	StockID   string    `gorm:"not null" json:"stock_id"`
	Quantity  uint      `gorm:"not null" json:"quantity"`
	Price     uint      `gorm:"not null" json:"price"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func CreateStockTransaction(s *[]StockTransaction) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Create(s).Error
}
