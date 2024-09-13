package routes

import (
	"awesomeProject2/views"

	"github.com/gofiber/fiber/v2"
)

func (route *Route) HelloRoute() {
	route.Group.Get("/hello/:name", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, Pg: route.Pg}
		return view.HelloView()
	})
}
