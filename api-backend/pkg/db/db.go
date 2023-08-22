package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func ConnectPostgreSqlDb(ctx context.Context) (*pgx.Conn, error) {
	url := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}
