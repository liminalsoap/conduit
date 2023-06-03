package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	Conn *pgx.Conn
}

func NewDb(url string) (*Postgres, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return &Postgres{Conn: conn}, nil
}
