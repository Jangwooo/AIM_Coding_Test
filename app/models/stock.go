package models

import (
	"time"

	"github.com/google/uuid"
)

type Stock struct {
	ID        uuid.UUID `gorm:"primaryKey"      json:"ID"`
	StockCode string    `gorm:"not null;unique" json:"stockCode"`
	StockName string    `gorm:"not null" json:"stockName"`
	Price     int64     `gorm:"not null" json:"price"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
