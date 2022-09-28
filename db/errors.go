package db

import "errors"

var (
	ErrUserNotExist        = errors.New("User does not exist in db")
	ErrBookNotExist        = errors.New("Book does not exist in db")
	ErrTransactionNotExist = errors.New("Transaction does not exist in db")
	ErrIDNotExist          = errors.New("ID does not exist in db")
	ErrAlreadyReturn       = errors.New("Book Already returned")
	//ErrUserTakenBook       = errors.New("Cannot delete the user as he has not returned the book")
)
