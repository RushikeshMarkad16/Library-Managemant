package transaction

import "errors"

var (
	errEmptyBookID    = errors.New("Book ID must be present")
	errEmptyUserID    = errors.New("User ID must be present")
	errNoTransactions = errors.New("No Transactions present")
	errAlreadyTaken   = errors.New("You cannot take this book again as it is issued to you already")
)
