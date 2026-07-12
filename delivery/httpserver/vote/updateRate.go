package votehandler

import (
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) UpdateRate(c *fiber.Ctx) error {
	id, errID := c.ParamsInt("rate_id")

	if errID != nil {
		return richerror.Out(
			richerror.New("updaterate.delivery").WithMessage("id not found").WithKind(richerror.KindInvalid),
			c)
	}

	var data voteparam.Vote

	errBody := c.BodyParser(&data)

	if errBody != nil {
		return richerror.Out(
			richerror.New("updaterate.delivery").WithMessage("error on parsing data").WithKind(richerror.KindInvalid),
			c)
	}

	errSvc := h.voteSvc.UpdateRate(id, data.Rate)

	if errSvc != nil {
		return richerror.Out(errSvc, c)
	}
	return c.JSON(fiber.Map{"status": "ok"})
}
