package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

func (u UserModel) Add(user User) error {
	_, err := u.DB.Exec(u.Ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u UserModel) Find(email, password string) (User, error) {
	user := User{}

	err := u.DB.QueryRow(u.Ctx, "SELECT id, email, password, created_at FROM users WHERE email=$1", email).Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, err
	}

	return user, nil
}
