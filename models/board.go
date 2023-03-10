package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Use `boards` table fields.
type Board struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UserId    string    `json:"user_id,omitempty"`
}

type CardsBoard struct {
	BoardName      string    `json:"board_name"`
	BoardCreatedAt time.Time `json:"created_at"`
	Cards          []Card    `json:"cards,omitempty"`
	// CardId         string    `json:"card_id"`
	// CardContent    string    `json:"card_content"`
	// CardCreatedAt  time.Time `json:"card_created_at"`
}

// Allow instanciate the database and context.
type BoardModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

// Insert a new board bind with username.
func (m BoardModel) Add(name, username string) error {
	_, err := m.DB.Exec(
		m.Ctx,
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

// Insert a new card into specified board and correct user.
func (m BoardModel) AddCard(content, username, boardId string) error {
	_, err := m.DB.Exec(
		m.Ctx,
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

// Select all boards by username.
func (m BoardModel) FindAll(username string) ([]Board, error) {
	rows, err := m.DB.Query(
		m.Ctx,
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

// Select one board by id and username and it's cards.
func (m BoardModel) FindOne(username, id string) (CardsBoard, error) {
	cardsBoard := CardsBoard{Cards: make([]Card, 0)}

	err := m.DB.QueryRow(
		m.Ctx,
		`SELECT name, created_at
			  FROM boards
			  WHERE id = $1`,
		id,
	).Scan(&cardsBoard.BoardName, &cardsBoard.BoardCreatedAt)
	if err != nil {
		return CardsBoard{}, err
	}

	rows, err := m.DB.Query(
		m.Ctx,
		`SELECT cards.id AS card_id, cards.content AS card_content, cards.created_at AS card_created_at
			  FROM boards
			  JOIN users ON boards.user_id = users.id
			  JOIN cards ON boards.id = cards.board_id
			  WHERE users.username = $1 AND boards.id = $2`,
		username, id,
	)
	if err != nil {
		return CardsBoard{}, err
	}
	defer rows.Close()

	for rows.Next() {
		c := Card{}
		err := rows.Scan(&c.Id, &c.Content, &c.CreatedAt)
		if err != nil {
			return CardsBoard{}, err
		}
		cardsBoard.Cards = append(cardsBoard.Cards, c)
	}

	return cardsBoard, nil
}

// Update a board name.
func (m BoardModel) Update(newName, id string) error {
	_, err := m.DB.Exec(
		m.Ctx,
		`UPDATE boards
			  SET name = $1
			  WHERE id = $2`,
		newName, id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete a board by id.
func (m BoardModel) Delete(id string) error {
	_, err := m.DB.Exec(
		m.Ctx,
		`DELETE FROM boards
			  WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
