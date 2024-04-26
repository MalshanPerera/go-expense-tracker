package database

import (
	"context"
	"fmt"
	"os"

	"github.com/MalshanPerera/go-expense-tracker/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func createConnStr(cfg *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
}

// Use pgxpool for concurrent connections
// if you have multiple threads working with a DB at the same time, you must use pgxpool
func Connect(cfg *config.Config) *pgxpool.Pool {
	connStr := createConnStr(cfg)

	var err error

	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = db.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")

	return db
}

func Close() {
	db.Close()
}
