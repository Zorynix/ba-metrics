package services

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/rs/zerolog/log"
)

type Clickhouse struct {
	db driver.Conn
}

func NewClickHouse(ctx context.Context) (*Clickhouse, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:19000"},
		Auth: clickhouse.Auth{
			Database: "ba_metrics",
			Username: "default",
			Password: "qwerty123",
		},
		DialTimeout:     time.Second * 30,
		Debug:           false,
		MaxOpenConns:    100,
		MaxIdleConns:    20,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		log.Fatal().Interface("unable to create clickhouse connection pool: %v", err).Msg("")
	}

	return &Clickhouse{db: conn}, nil
}

func (ch *Clickhouse) Ping(ctx context.Context) error {
	return ch.db.Ping(ctx)
}

func (ch *Clickhouse) Close() {
	ch.db.Close()
}
