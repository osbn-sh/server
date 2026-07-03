package homehandler

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func (h Handler) SetRoutes(e *fiber.App) {
	e.Get("/", h.IndexPage)

	e.Get("/swagger/*", fiberSwagger.WrapHandler)

	e.Get("/doc", h.InitScalarDoc)

	e.Get("/openapi.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

}
