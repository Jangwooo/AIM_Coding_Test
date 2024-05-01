package models

import (
	"github.com/google/uuid"
	"time"
)

type RiskType string

const (
	RiskTypeHigh RiskType = "high"
	RiskTypeLow  RiskType = "low"
)

type Portfolio struct {
	ID             uuid.UUID        `gorm:"primaryKey" json:"id"`
	UserID         string           `gorm:"not null" json:"user_id"`
	RiskType       RiskType         `gorm:"not null" json:"risk_type"` // enum: 위험, 안전
	CreatedAt      time.Time        `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	PortfolioItems *[]PortfolioItem `gorm:"foreignKey:PortfolioID" json:"portfolio_items,omitempty"`
}
