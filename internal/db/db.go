package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"example.com/go-bank/migrations"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib" // Required to bridge pgx to goose
	"github.com/pressly/goose/v3"
)

var pool *pgxpool.Pool


func InitDB() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	var err error
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/gobanking"
	}
	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	// Test connection
    if err := pool.Ping(ctx); err != nil {
		return nil, err
    }
	fmt.Println("Connected to database successfully.")

	// BRIDGE (only needed for goose migrations)
	dbSQL := stdlib.OpenDBFromPool(pool)
	defer dbSQL.Close()


	// set goose to use the embedded migrations (SetBaseFS)
	// set db `dialect` for goose to "postgres" (SetDialect)
	// run migrations  (goose.Up)

	
	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}
	if err := goose.Up(dbSQL, "."); err != nil {
		return nil, err
	}

	return pool, nil
}
