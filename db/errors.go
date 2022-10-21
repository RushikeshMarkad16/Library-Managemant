package db

import "errors"

var (
	ErrUserNotExist        = errors.New("Invalid user id")
	ErrBookNotExist        = errors.New("Invalid book id")
	ErrTransactionNotExist = errors.New("Invalid transaction id")
	ErrIDNotExist          = errors.New("Invalid user id")
	ErrAlreadyReturn       = errors.New("Book Already returned")
)
