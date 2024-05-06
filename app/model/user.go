package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"github.com/Jangwooo/AIM_Coding_Test/pkg/utils"
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)

type User struct {
	ID               string              `gorm:"primaryKey" json:"id"`
	Username         string              `gorm:"not null" json:"username"`
	Password         string              `gorm:"not null" json:"password"`
	CreatedAt        time.Time           `gorm:"not null" json:"created_at"`
	UpdatedAt        time.Time           `json:"updated_at"`
	LoginLogs        *[]LoginLog         `gorm:"foreignKey:UserID" json:"login_logs,omitempty"`
	Portfolios       *[]Portfolio        `gorm:"foreignKey:UserID" json:"portfolios,omitempty"`
	Account          *Account            `gorm:"foreignKey:UserID" json:"account,omitempty"`
	StockTransaction *[]StockTransaction `gorm:"foreignKey:UserID" json:"stock_transaction,omitempty"`
}

func (u *User) CreateUser() error {
	u.Password = utils.GeneratePassword(u.Password)

	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	err = db.Create(u).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return pkg.ErrDuplicateUser
			}
		}
	}

	return err
}

func (u *User) GetUserByID(id string) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Where("ID = ?", id).First(&u).Error
}
