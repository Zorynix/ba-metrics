package routes

import (
	"awesomeProject2/services"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type RouterHead struct {
	Pg         *services.Postgres
	Clickhouse *services.Clickhouse
	Addr       *string
}

type Router struct {
	Router     *fiber.App
	Pg         *services.Postgres
	Clickhouse *services.Clickhouse
}

type Route struct {
	Group      fiber.Router
	Pg         *services.Postgres
	Clickhouse *services.Clickhouse
}

func Routes(addr *string) {
	postgres, err := services.NewPG(context.Background())
	if err != nil {
		log.Fatal().Err(err)
	}

	clickhouse, err := services.NewClickHouse(context.Background())
	if err != nil {
		log.Fatal().Err(err)
	}

	router := fiber.New()

	route := Router{Router: router, Pg: postgres, Clickhouse: clickhouse}

	route.V1Routes()
	route.MetricsRoutes()

	if err := router.Listen(*addr); err != nil {
		log.Fatal().Err(err).Msg("Can not start http server")
	}
}
