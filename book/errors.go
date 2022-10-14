package book

import "errors"

var (
	errEmptyID                 = errors.New("Book ID must be present")
	errEmptyName               = errors.New("Book name cannot be empty")
	errEmptyAuthor             = errors.New("Author name must be present")
	errZeroCopies              = errors.New("Copies cannot be zero while creation of book")
	errInvalidPrice            = errors.New("Price should be greater than zero")
	errInvalidStatus           = errors.New("Invalid Status of book")
	errInvalidAvailableCopies  = errors.New("Available copies cannot be greater than total copies")
	errInvalidTotalCopies      = errors.New("Total copies must be integer")
	err1InvalidAvailableCopies = errors.New("Available copies must be integer")

	errNoBooks  = errors.New("No book present")
	errNoBookId = errors.New("Book is not present")
)
