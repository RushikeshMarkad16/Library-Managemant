package transaction

import "errors"

var (
	errEmptyBookID = errors.New("Book ID must be present")
	errEmptyUserID = errors.New("User ID must be present")
	//errInvalidReturnDate = errors.New("Return Date cannot be less than issue date")
	errNoTransactions = errors.New("No Transactions present")
	errAlreadyTaken   = errors.New("You cannot take this book again as it is issued to you already")
)
