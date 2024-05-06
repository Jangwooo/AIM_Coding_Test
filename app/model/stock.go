package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)

type Stock struct {
	ID               string              `gorm:"primaryKey" json:"stock_id"`
	StockName        string              `gorm:"not null" json:"stock_name"`
	Price            uint                `gorm:"not null" json:"price"`
	CreatedAt        time.Time           `gorm:"not null" json:"createdAt"`
	UpdatedAt        time.Time           `json:"updatedAt"`
	StockTransaction *[]StockTransaction `gorm:"foreignKey:stock_id" json:"stock_transaction,omitempty"`
}

func (s *Stock) CreateStock() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	err = db.Create(&s).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return pkg.ErrDuplicateStockID
			}
		}
	}

	return err
}

func (s *Stock) GetStockByStockCode() (*Stock, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	err = db.Where("id = ?", s.ID).First(&s).Error
	return s, err
}

func (s *Stock) UpdateStock() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Updates(s).Error
}

func (s *Stock) DeleteStock() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Delete(&s).Error
}

func GetStockList() ([]Stock, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var stocks []Stock
	err = db.Find(&stocks).Error
	return stocks, err
}
