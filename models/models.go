package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

func (u UserModel) Add(user User) error {
	_, err := u.DB.Exec(
		u.Ctx,
		`INSERT INTO users (username, email, password)
			  VALUES ($1, $2, $3)`,
		user.Username, user.Email, user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

// FindOne
func (u UserModel) Find(email, password string) (User, error) {
	user := User{}

	err := u.DB.QueryRow(
		u.Ctx,
		`SELECT id, username, email, password, created_at
		      FROM users
			  WHERE email = $1`,
		email,
	).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, err
	}

	return user, nil
}
