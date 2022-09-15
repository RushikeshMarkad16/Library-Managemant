package book

import "errors"

var (
	errEmptyID                = errors.New("Book ID must be present")
	errEmptyName              = errors.New("Book name must be present")
	errEmptyAuthor            = errors.New("Author name must be present")
	errZeroCopies             = errors.New("Copies cannot be zero while creation of book")
	errInvalidPrice           = errors.New("Price should be greater than zero")
	errInvalidStatus          = errors.New("Invalid Status of book")
	errInvalidAvailableCopies = errors.New("Available copies cannot be greater than total copies")

	errNoBooks  = errors.New("No book present")
	errNoBookId = errors.New("Book is not present")
)
