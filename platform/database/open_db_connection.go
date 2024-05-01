package database

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"gorm.io/gorm"
)

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*gorm.DB, error) {
	// Define Database connection variables.
	var (
		err error
		db  *gorm.DB
	)

	// Get DB_TYPE value from .app.env file

	// Define a new Database connection with right DB type.
	db, err = MysqlConnection()

	return db, err
}

func InitDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transaction{},
		&models.Portfolio{},
		&models.PortfolioItem{},
		&models.Stock{},
		&models.LoginLog{})

	return err
}
