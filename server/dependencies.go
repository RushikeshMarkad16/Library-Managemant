package server

import (
	"github.com/RushikeshMarkad16/Library-Managemant/app"
	"github.com/RushikeshMarkad16/Library-Managemant/book"
	"github.com/RushikeshMarkad16/Library-Managemant/db"
	"github.com/RushikeshMarkad16/Library-Managemant/user"
)

type dependencies struct {
	UserService user.Service
	BookService book.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := user.NewService(dbStore, logger)
	bookService := book.NewService(dbStore, logger)

	return dependencies{
		UserService: userService,
		BookService: bookService,
	}, nil
}
