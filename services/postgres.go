package services

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic().Msg("No .env file found")
	}
}

func NewPG(ctx context.Context) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, os.Getenv("POSTGRES_DSN"))
		if err != nil {
			log.Fatal().Interface("unable to create connection pool: %v", err)
		}

		pgInstance = &Postgres{db}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}
