package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	connIdleTimeout = 10 * time.Minute
	connMaxLifetime = 1 * time.Hour
	maxIdleConn     = 10
)

func ConnectPostgreSqlDb(ctx context.Context) (*sqlx.DB, error) {
	url := os.Getenv("DATABASE_URL")

	c, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	connStr := stdlib.RegisterConnConfig(c)

	conn, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatalln(err)
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	conn.SetMaxIdleConns(maxIdleConn)
	conn.SetConnMaxIdleTime(connIdleTimeout)
	conn.SetConnMaxLifetime(connMaxLifetime)

	return conn, nil
}
