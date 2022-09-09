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
	ID         int    `db:"id"`
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

func (s *store) CreateCategory(ctx context.Context, user *User) (err error) {
	now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			user.First_name,
			now,
			now,
		)
		return err
	})
}
