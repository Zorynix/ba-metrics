package views

import (
	"awesomeProject2/services"

	"github.com/gofiber/fiber/v2"
)

type View struct {
	Ctx        *fiber.Ctx
	Pg         *services.Postgres
	Clickhouse *services.Clickhouse
}
