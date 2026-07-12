package votehandler

import (
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) DeleteOption(c *fiber.Ctx) error {

	id, errID := c.ParamsInt("option_id")

	if errID != nil {
		return richerror.Out(
			richerror.New("deleteOption.delivery").WithMessage("id not found").WithKind(richerror.KindInvalid),
			c)
	}

	errSvc := h.voteSvc.RemoveOption(id)

	if errSvc != nil {
		return richerror.Out(
			errSvc,
			c)
	}

	return c.JSON(fiber.Map{
		"success": "true",
	})
}
