package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"time"
)

type Portfolio struct {
	ID             string           `gorm:"primaryKey" json:"id"` // UUID
	UserID         string           `gorm:"not null" json:"user_id"`
	RiskType       string           `gorm:"not null" json:"risk_type"` // enum: 위험, 안전
	CreatedAt      time.Time        `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	PortfolioItems *[]PortfolioItem `gorm:"foreignKey:portfolio_id" json:"portfolio_items,omitempty"`
}

func (p *Portfolio) CreatePortfolio() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Create(p).Error
}

func (p *Portfolio) GetPortfolioByID() (*Portfolio, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	err = db.Model(p).Preload("PortfolioItems").First(p).Error
	return p, err
}

func (p *Portfolio) GetPortfolioList() ([]Portfolio, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var portfolios []Portfolio

	err = db.Find(&portfolios, "user_id = ?", p.UserID).Order("created_at desc").Error

	return portfolios, err
}
