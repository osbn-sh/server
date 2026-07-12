package votehandler

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) RateUniv(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"success": "true",
	})
}
