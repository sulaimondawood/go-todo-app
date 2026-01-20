package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(connectionStr string) (*pgxpool.Pool, error) {

	ctx := context.Background()

	var config *pgxpool.Config
	var err error

	config, err = pgxpool.ParseConfig(connectionStr)

	if err != nil {
		log.Panicf("Unable to parse Database URL: %v", err)
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Printf("Unable to crate db connection pool: %v", err)
		return nil, err
	}

}
