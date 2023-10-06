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
			($1, $2)
		ON CONFLICT (name)
		DO UPDATE SET count = users.count + 1
		RETURNING users.name, users.count`, name)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.User])

	return user, nil
}
