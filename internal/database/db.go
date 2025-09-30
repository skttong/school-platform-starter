package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"school/internal/config"
)

func Connect(cfg *config.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		log.Panicf("db connect error: %v", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		log.Panicf("db ping error: %v", err)
	}
	return pool
}
