package views

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (view *View) HelloView() error {
	name := view.Ctx.Params("name")

	user, err := view.Pg.SayHello(name)
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadRequest)
	}

	var payload string
	if user.Count > 0 {
		payload = fmt.Sprintf("Hello against, %s", user.Name)
	} else {
		payload = fmt.Sprintf("Hello %s", user.Name)
	}

	return view.Ctx.SendString(payload)
}
