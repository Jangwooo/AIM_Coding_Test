package database

import (
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"

	"github.com/Jangwooo/AIM_Coding_Test/pkg/utils"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // load driver for Mysql
)

// GetConnection func for connection to Mysql database.
func GetConnection() (*gorm.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build Mysql connection URL.
	mysqlConnURL := utils.ConnectionURLBuilder("mysql")

	// Define database connection for Mysql.
	db, err := sqlx.Connect("mysql", mysqlConnURL)
	if err != nil {
		return nil, errors.Wrap(err, pkg.ErrDatabaseConnectionFailed.Error())
	}

	// Set database connection settings:
	// 	- SetMaxOpenConns: the default is 0 (unlimited)
	// 	- SetMaxIdleConns: defaultMaxIdleConns = 2
	// 	- SetConnMaxLifetime: 0, connections are reused forever
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	orm, err := gorm.Open(mysql.New(mysql.Config{Conn: db, DefaultStringSize: 191}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, errors.Wrap(err, pkg.ErrDatabaseConnectionFailed.Error())
	}

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, errors.Wrap(err, pkg.ErrDatabaseOperationFailed.Error())
	}

	return orm, nil
}
