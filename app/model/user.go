package model

import (
	"github.com/Jangwooo/AIM_Coding_Test/pkg/utils"
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"time"
)

type User struct {
	ID         string       `gorm:"primaryKey" json:"id"`
	Username   string       `gorm:"not null" json:"username"`
	Password   string       `gorm:"not null" json:"password"`
	CreatedAt  time.Time    `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	LoginLogs  *[]LoginLog  `gorm:"foreignKey:UserID" json:"login_logs,omitempty"`
	Portfolios *[]Portfolio `gorm:"foreignKey:UserID" json:"portfolios,omitempty"`
	Account    *Account     `gorm:"foreignKey:UserID" json:"account,omitempty"`
}

func (u *User) CreateUser(id, pwd, username string) error {
	u.ID = id
	u.Password = utils.GeneratePassword(pwd)
	u.Username = username
	u.CreatedAt = time.Now()

	db := database.OpenDBConnection()

	return db.Create(u).Error
}

func (u *User) GetUserByID(id string) (*User, error) {
	db := database.OpenDBConnection()

	err := db.Where("ID = ?", id).First(&u).Error

	return u, err
}
