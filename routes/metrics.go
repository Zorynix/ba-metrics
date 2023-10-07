package routes

import (
	"awesomeProject2/views"

	"github.com/gofiber/fiber/v2"
)

func (route *Router) MetricsRoutes() {
	route.Router.Get("/to/:id", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, Pg: route.Pg, Clickhouse: route.Clickhouse}
		return view.MetricsView()
	})
}
