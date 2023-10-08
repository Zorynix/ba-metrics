package views

import (
	"awesomeProject2/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/rs/zerolog/log"
)

func (view *View) MetricsView() error {
	id, err := uuid.Parse(view.Ctx.Params("id"))
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.ErrBadRequest.Code)
	}
	link, err := view.Pg.TakeLink(id)
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.ErrBadRequest.Code)
	}

	fingerprints := services.TakeFingerprints(view.Ctx)
	err = view.Clickhouse.RecordMetrics(id, fingerprints)
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.ErrBadGateway.Code)
	}

	return view.Ctx.Redirect(link.Url)
}
