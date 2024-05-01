package models

import (
	"github.com/google/uuid"
	"time"
)

type PortfolioItem struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	PortfolioID string    `gorm:"not null" json:"portfolio_id"`
	StockCode   string    `gorm:"not null" json:"stock_code"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
