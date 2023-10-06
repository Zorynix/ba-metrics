package routes

import (
	"awesomeProject2/services"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type RouterHead struct {
	Pg   *services.Postgres
	Addr *string
}

type Router struct {
	Router *fiber.App
	Pg     *services.Postgres
}

type Route struct {
	Group fiber.Router
	Pg    *services.Postgres
}

func Routes(addr *string) {
	db, err := services.NewPG(context.Background())
	if err != nil {
		log.Fatal().Err(err)
	}

	router := fiber.New()

	route := Router{Router: router, Pg: db}

	route.V1Routes()

	if err := router.Listen(*addr); err != nil {
		log.Fatal().Err(err).Msg("Can not start http server")
	}
}
