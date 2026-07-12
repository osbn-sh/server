package votehandler

import (
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) UpdateOption(c *fiber.Ctx) error {
	id, errID := c.ParamsInt("option_id")

	if errID != nil {
		return richerror.Out(
			richerror.New("updaterate.delivery").WithMessage("id not found").WithKind(richerror.KindInvalid),
			c)
	}

	var data voteparam.Option

	errBody := c.BodyParser(&data)

	if errBody != nil {
		return richerror.Out(
			richerror.New("updaterate.delivery").WithMessage("error on parsing data").WithKind(richerror.KindInvalid),
			c)
	}

	errSvc := h.voteSvc.UpdateOption(id, data)

	return richerror.Out(errSvc, c)
}
