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

type Board struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UserId    string    `json:"user_id,omitempty"`
}

type CardsBoard struct {
	BoardId        string    `json:"board_id"`
	BoardName      string    `json:"board_name"`
	BoardCreatedAt time.Time `json:"board_created_at"`
	CardId         string    `json:"card_id"`
	CardContent    string    `json:"card_content"`
	CardCreatedAt  time.Time `json:"card_created_at"`
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

// FindOne
func (u UserModel) Find(email, password string) (User, error) {
	user := User{}

	err := u.DB.QueryRow(
		u.Ctx,
		"SELECT id, email, password, created_at FROM users WHERE email = $1",
		email,
	).Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, err
	}

	return user, nil
}

type BoardModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

func (b BoardModel) Add(name, email string) error {
	_, err := b.DB.Exec(
		b.Ctx,
		`INSERT INTO boards (name, user_id)
		      SELECT $1, id
			  FROM users
			  WHERE email = $2`,
		name, email,
	)
	if err != nil {
		return err
	}

	return nil
}

func (b BoardModel) FindAll(username string) ([]Board, error) {
	rows, err := b.DB.Query(
		b.Ctx,
		`SELECT boards.id, boards.name, boards.created_at
			  FROM boards
			  JOIN users ON boards.user_id = users.id
			  WHERE users.username = $1`,
		username,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	boards := make([]Board, 0)

	for rows.Next() {
		b := Board{}
		err := rows.Scan(&b.Id, &b.Name, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		boards = append(boards, b)
	}

	return boards, nil
}

func (b BoardModel) FindOne(email, boardId string) ([]CardsBoard, error) {
	rows, err := b.DB.Query(
		b.Ctx,
		`SELECT boards.id AS board_id, boards.name AS board_name, boards.created_at AS board_created_at, cards.id AS card_id, cards.content AS card_content, cards.created_at AS card_created_at
			  FROM boards
			  JOIN users ON boards.user_id = users.id
			  JOIN cards ON boards.id = cards.board_id
			  WHERE users.email = $1 AND boards.id = $2`,
		email, boardId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cardsBoard := make([]CardsBoard, 0)

	for rows.Next() {
		c := CardsBoard{}
		err := rows.Scan(&c.BoardId, &c.BoardName, &c.BoardCreatedAt, &c.CardId, &c.CardContent, &c.CardCreatedAt)
		if err != nil {
			return nil, err
		}
		cardsBoard = append(cardsBoard, c)
	}

	return cardsBoard, nil
}

// TODO: Create body to update a board name
func (b BoardModel) Update(board Board, userId string) error {
	return nil
}
