package routes

import (
	"awesomeProject2/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (route *Router) MetricsRoutes() {
	route.Router.Get("/to/:id", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, Pg: route.Pg, Clickhouse: route.Clickhouse}
		return view.MetricsView()
	})

	route.Router.Get("/metrics", monitor.New())
}
