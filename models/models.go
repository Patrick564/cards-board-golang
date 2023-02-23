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

type Card struct {
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UserId    string    `json:"user_id,omitempty"`
	BoardId   string    `json:"board_id"`
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

func (b BoardModel) Add(name, username string) error {
	_, err := b.DB.Exec(
		b.Ctx,
		`INSERT INTO boards (name, user_id)
		      SELECT $1, id
			  FROM users
			  WHERE username = $2`,
		name, username,
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

func (b BoardModel) FindOne(username, boardId string) ([]CardsBoard, error) {
	rows, err := b.DB.Query(
		b.Ctx,
		`SELECT boards.id AS board_id, boards.name AS board_name, boards.created_at AS board_created_at, cards.id AS card_id, cards.content AS card_content, cards.created_at AS card_created_at
			  FROM boards
			  JOIN users ON boards.user_id = users.id
			  JOIN cards ON boards.id = cards.board_id
			  WHERE users.username = $1 AND boards.id = $2`,
		username, boardId,
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
func (b BoardModel) Update(newName, id string) error {
	_, err := b.DB.Exec(
		b.Ctx,
		"UPDATE boards SET name = $1 WHERE id = $2",
		newName, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (b BoardModel) Delete(id string) error {
	_, err := b.DB.Exec(
		b.Ctx,
		"DELETE FROM boards WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

type CardModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

func (c CardModel) Add(username, boardId, content string) error {
	_, err := c.DB.Exec(
		c.Ctx,
		`INSERT INTO cards (content, user_id, board_id)
		      SELECT $1, users.id, $2
              FROM users
              WHERE username = $3`,
		content, boardId, username,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c CardModel) FindAllByUsername(username string) ([]Card, error) {
	rows, err := c.DB.Query(
		c.Ctx,
		`SELECT cards.id, cards.content, cards.created_at, cards.board_id
			  FROM cards
			  JOIN users ON cards.user_id  = users.id
			  WHERE users.username = $1`,
		username,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]Card, 0)

	for rows.Next() {
		c := Card{}
		err := rows.Scan(&c.Id, &c.Content, &c.CreatedAt, &c.BoardId)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}

	return cards, nil
}

func (c CardModel) FindAllByBoardId(boardId string) ([]Card, error) {
	return nil, nil
}

func (c CardModel) FindOne(id string) (Card, error) {
	return Card{}, nil
}

func (c CardModel) Update(id string) error {
	return nil
}

func (c CardModel) Delete(id string) error {
	return nil
}
