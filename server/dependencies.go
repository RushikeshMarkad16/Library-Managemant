package server

import (
	"github.com/RushikeshMarkad16/Library-Managemant/app"
	"github.com/RushikeshMarkad16/Library-Managemant/book"
	"github.com/RushikeshMarkad16/Library-Managemant/db"
	"github.com/RushikeshMarkad16/Library-Managemant/transaction"
	"github.com/RushikeshMarkad16/Library-Managemant/user"
)

type dependencies struct {
	UserService        user.Service
	BookService        book.Service
	TransactionService transaction.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := user.NewService(dbStore, logger)
	bookService := book.NewService(dbStore, logger)
	transactionService := transaction.NewService(dbStore, logger)

	return dependencies{
		UserService:        userService,
		BookService:        bookService,
		TransactionService: transactionService,
	}, nil
}
