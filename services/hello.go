package services

import (
	"awesomeProject2/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) SayHello(name string) (models.User, error) {
	rows, err := pg.db.Query(context.Background(), `
		INSERT INTO
			service_template_schema.users (name, count)
		VALUES
			($1,1)
		ON CONFLICT (name)
		DO UPDATE SET count = users.count + 1
		RETURNING users.name, users.count`, name)
	if err != nil {
		return models.User{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.User])
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

//localhost:8000/to/uuid
//
