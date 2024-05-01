package database

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

// OpenDBConnection func for opening database connection.
func OpenDBConnection() *gorm.DB {
	// Define Database connection variables.
	var (
		err error
		db  *gorm.DB
	)

	// Get DB_TYPE value from .app.env file

	// Define a new Database connection with right DB type.
	db, err = MysqlConnection()
	if err != nil {
		log.Fatal(fmt.Errorf("open db error: %v", err))
	}

	return db
}
