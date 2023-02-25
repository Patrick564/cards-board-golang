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

type Card struct {
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UserId    string    `json:"user_id,omitempty"`
	BoardId   string    `json:"board_id"`
}

type CardsBoard struct {
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
	rows, err := c.DB.Query(
		c.Ctx,
		`SELECT id, content, created_at, user_id
			  FROM cards
			  WHERE board_id = $1`,
		boardId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]Card, 0)

	for rows.Next() {
		c := Card{}
		err := rows.Scan(&c.Id, &c.Content, &c.CreatedAt, &c.UserId)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}

	return cards, nil
}

func (c CardModel) FindOne(username, id string) (Card, error) {
	card := Card{}

	err := c.DB.QueryRow(
		c.Ctx,
		`SELECT id, content, board_id
			  FROM cards
			  WHERE id = $1`,
		id,
	).Scan(&card.Id, &card.Content, &card.BoardId)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

func (c CardModel) Update(content, id string) error {
	_, err := c.DB.Exec(
		c.Ctx,
		"UPDATE cards SET content = $1 WHERE id = $2",
		content, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c CardModel) Delete(id string) error {
	_, err := c.DB.Exec(
		c.Ctx,
		"DELETE FROM cards WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
