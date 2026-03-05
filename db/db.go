package db

import (
	"context"
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib" // Required to bridge pgx to goose
	"github.com/pressly/goose/v3"
)

var Pool *pgxpool.Pool

//go:embed migrations/*.sql
var embedMigrations embed.FS //for migrations

func InitDB() {
	var err error
	dsn := "postgres://user:password@localhost:5432/gobanking"
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := Pool.Ping(ctx); err != nil {
        log.Fatalf("DB connection failed: %v", err)
    }
	fmt.Println("Connected to database successfully.")

	// BRIDGE
	DB := stdlib.OpenDBFromPool(Pool)
	defer DB.Close()


	// Set goose to use the embedded migrations
	goose.SetBaseFS(embedMigrations)

	// set db `dialect` for goose
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v\n", err)
	}

	// migrations (creating database tables)
	if err := goose.Up(DB, "migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v\n", err)
	}

}
