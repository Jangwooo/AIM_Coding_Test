package pkg

import (
	"errors"
)

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrPasswordNotMatch         = errors.New("password not match")
	ErrDuplicateUser            = errors.New("duplicate user")
	ErrDatabaseConnectionFailed = errors.New("database connection failed")
	ErrDatabaseOperationFailed  = errors.New("database operation failed")
	ErrSessionExpired           = errors.New("session expired. please login again")
	ErrForbiddenResourceAccess  = errors.New("forbidden to access this resource")
	ErrOnlyAdminCanAccess       = errors.New("only admin can access this resource")
	ErrDuplicateAccountID       = errors.New("duplicate account id")
	ErrRecordNotFound           = errors.New("record not found")
	ErrBalanceNotEnough         = errors.New("balance not enough")
	ErrDuplicateStockID         = errors.New("duplicate stock id")
)
