package db

import (
	"context"
	"database/sql"
)

const (
	createBookQuery = `INSERT INTO book ( 
    id,name,author,price,total_copies,status,available_copies)
    VALUES(?, ?, ?, ?, ?, ?, ?)`

	listBooksQuery      = `SELECT * FROM book`
	findBookByIDQuery   = `SELECT * FROM book WHERE id = ?`
	deleteBookByIDQuery = `DELETE FROM book WHERE id = ?`
	updateBookQuery     = `UPDATE book SET name=?, author=?, price=?, total_copies=?, status=?, available_copies=? WHERE id=? `
)

type Book struct {
	ID              string `db:"id"`
	Name            string `db:"name"`
	Author          string `db:"author"`
	Price           int    `db:"price"`
	TotalCopies     int    `db:"total_copies"`
	Status          string `db:"status"`
	AvailableCopies int    `db:"available_copies"`
}

func (s *store) CreateBook(ctx context.Context, book *Book) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createBookQuery,
			book.ID,
			book.Name,
			book.Author,
			book.Price,
			book.TotalCopies,
			book.Status,
			book.AvailableCopies,
		)
		return err
	})
}

func (s *store) ListBooks(ctx context.Context) (books []Book, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &books, listBooksQuery)
	})
	if err == sql.ErrNoRows {
		return books, ErrBookNotExist
	}
	return
}

func (s *store) FindBookByID(ctx context.Context, id string) (book Book, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &book, findBookByIDQuery, id)
	})
	if err == sql.ErrNoRows {
		return book, ErrBookNotExist
	}
	return
}

func (s *store) DeleteBookByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteBookByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrBookNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateBook(ctx context.Context, book *Book) (err error) {
	//now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateBookQuery,
			book.Name,
			book.Author,
			book.Price,
			book.TotalCopies,
			book.Status,
			book.AvailableCopies,
			book.ID,
		)
		return err
	})
}
