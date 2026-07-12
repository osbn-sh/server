package votehandler

import (
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetRate(c *fiber.Ctx) error {

	var data voteparam.Vote

	errBody := c.BodyParser(&data)

	if errBody != nil {
		return richerror.Out(
			richerror.New("addOption.delivery").WithMessage("error on parsing data").WithKind(richerror.KindInvalid),
			c)
	}

	SvcData, err := h.voteSvc.Get(data)

	if err != nil {
		return richerror.Out(err, c)
	}

	return c.JSON(SvcData)

}
