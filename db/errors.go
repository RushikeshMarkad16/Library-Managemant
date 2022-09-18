package db

import "errors"

var (
	ErrUserNotExist        = errors.New("User does not exist in db")
	ErrBookNotExist        = errors.New("Book does not exist in db")
	ErrTransactionNotExist = errors.New("Transaction does not exist in db")
	ErrUserTakenBook       = errors.New("Cannot delete the user as he has not returned the book")
)
