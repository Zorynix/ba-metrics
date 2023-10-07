package services

import (
	"awesomeProject2/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) TakeLink(id uuid.UUID) (models.Link, error) {
	rows, err := pg.db.Query(context.Background(), `
		SELECT id, url FROM ba_schema.links 
		WHERE id = $1`, id)
	if err != nil {
		return models.Link{}, err
	}

	link, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Link])
	if err != nil {
		return models.Link{}, err
	}

	return link, nil
}

func TakeFingerprints(c *fiber.Ctx) models.Fingerprints {
	ip := c.IP()
	userAgent := c.Get("User-Agent")
	referer := c.Get("Referer")

	return models.Fingerprints{Ip: ip, UserAgent: userAgent, Referer: referer}
}

func (ch *Clickhouse) RecordMetrics(linkId uuid.UUID, fingerprints models.Fingerprints) error {
	err := ch.db.Exec(context.Background(), `
	INSERT INTO 
		ba_metrics.metrics (created_at, link_id, ip, user_agent, referer)
		VALUES
			(now(), $1, $2, $3, $4)
	`, linkId, fingerprints.Ip, fingerprints.UserAgent, fingerprints.Referer)
	if err != nil {
		return err
	}

	return nil
}
