package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Conn struct {
	Pool *pgxpool.Pool
}

func New(dsn string) (*Conn, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("error creating db connection pool %w", err)
	}
	return &Conn{Pool: pool}, nil
}

func (c *Conn) Close() {
	if c.Pool != nil {
		c.Pool.Close()
	}
}
