package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(host string, port int, dbName string, userName string, password string) (*Storage, error) {
	const op = "storage.postgesql.New"

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", userName, password, dbName, host, port)

	ctx, cansel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cansel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: cant't connect to database: %w", op, err)
	}

	ctxPing, cancelPing := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelPing()
	if err := pool.Ping(ctxPing); err != nil {
		return nil, fmt.Errorf("%s: can't ping database: %w")
	}

	return &Storage{pool: pool}, nil
}
