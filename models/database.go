package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Conn struct {
	DB *pgx.Conn
}

func Connect() (Conn, error) {
	dsn := ""
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return Conn{}, err
	}

	return Conn{DB: conn}, nil
}
