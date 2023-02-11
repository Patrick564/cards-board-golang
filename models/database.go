package models

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(ctx context.Context, url string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(ctx, url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
