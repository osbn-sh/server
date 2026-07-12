package votehandler

import (
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) AddOption(c *fiber.Ctx) error {

	var data voteparam.Option

	errBody := c.BodyParser(&data)

	if errBody != nil {
		return richerror.Out(
			richerror.New("addOption.delivery").WithMessage("error on parsing data").WithKind(richerror.KindInvalid),
			c)
	}
	
	errSvc := h.voteSvc.AddOption(data)

	if errSvc != nil {
		return richerror.Out(
			errSvc,
			c)
	}

	return c.JSON(fiber.Map{
		"success": "true",
	})
}
