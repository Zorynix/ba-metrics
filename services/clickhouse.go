package services

import (
	"context"
	"crypto/tls"
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
		Addr: []string{"127.0.0.1:9009"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout:          time.Second * 30,
		Debug:                true,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
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
