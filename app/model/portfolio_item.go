package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"time"
)

type PortfolioItem struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	PortfolioID string    `gorm:"not null" json:"portfolio_id"`
	StockCode   string    `gorm:"not null" json:"stock_code"`
	Quantity    uint      `gorm:"not null" json:"quantity"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreatePortfolioItem(item *[]PortfolioItem) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Create(item).Error
}
