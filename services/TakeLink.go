package services

import (
	"awesomeProject2/models"
	"context"

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
