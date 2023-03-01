package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Card struct {
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UserId    string    `json:"user_id,omitempty"`
	BoardId   string    `json:"board_id,omitempty"`
}

type CardModel struct {
	DB  *pgxpool.Pool
	Ctx context.Context
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
