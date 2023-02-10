package models

import (
	"time"

	"github.com/jackc/pgx/v4"
)

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserModel struct {
	DB *pgx.Conn
}

func (u UserModel) Register() (User, error) {
	return User{}, nil
}
