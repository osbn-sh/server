package userhandler

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) IsLogin(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"success": "true",
	})
}
