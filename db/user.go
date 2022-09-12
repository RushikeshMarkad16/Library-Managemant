package db

import (
	"context"
	"database/sql"
	"time"
)

const (
	createUserQuery = `INSERT INTO user ( 
    id,first_name,last_name,gender,dob,address,email,password,mob_no,role)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	listUsersQuery      = `SELECT * FROM user`
	findUserByIDQuery   = `SELECT * FROM user WHERE id = $1`
	deleteUserByIDQuery = `DELETE FROM user WHERE id = $1`
	updateUserQuery     = `UPDATE user SET name = $1, updated_at = $2 where id = $3`
)

type User struct {
	ID         string `db:"id"`
	First_name string `db:"first_name"`
	Last_name  string `db:"last_name"`
	Gender     string `db:"gender"`
	DOB        string `db:"dob"`
	Address    string `db:"address"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Mob_no     int    `db:"mob_no"`
	Role       string `db:"role"`
}

func (s *store) CreateUser(ctx context.Context, user *User) (err error) {
	//now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			user.ID,
			user.First_name,
			user.Last_name,
			user.Gender,
			user.DOB,
			user.Address,
			user.Email,
			user.Password,
			user.Mob_no,
			user.Role,
			// now,
			// now,
		)
		return err
	})
}

func (s *store) ListUsers(ctx context.Context) (users []User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &users, listUsersQuery)
	})
	if err == sql.ErrNoRows {
		return users, ErrUserNotExist
	}
	return
}

func (s *store) FindUserByID(ctx context.Context, id string) (user User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByIDQuery, id)
	})
	if err == sql.ErrNoRows {
		return user, ErrUserNotExist
	}
	return
}

func (s *store) DeleteUserByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteUserByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrUserNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateUser(ctx context.Context, user *User) (err error) {
	now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateUserQuery,
			user.First_name,
			now,
			user.ID,
		)
		return err
	})
}
