package userhandler

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Login(c *fiber.Ctx) error {

	//h.userSvc.IsExist()

	return c.JSON(fiber.Map{
		"success": true,
	})
}
