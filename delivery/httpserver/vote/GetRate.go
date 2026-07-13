package votehandler

import (
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetRate(c *fiber.Ctx) error {

	entity := c.Params("entity")

	if entity == "" {
		return richerror.Out(
			richerror.New("voteget.delivery").WithMessage("entity not found").WithKind(richerror.KindInvalid),
			c)
	}

	slug, errIDSlug := c.ParamsInt("slug")

	if errIDSlug != nil {
		return richerror.Out(
			richerror.New("voteget.delivery").WithMessage("slug not found").WithKind(richerror.KindInvalid),
			c)
	}

	data := voteparam.Vote{
		Target:   entity,
		TargetID: slug,
	}

	SvcData, err := h.voteSvc.Get(data)

	if err != nil {
		return richerror.Out(err, c)
	}

	return c.JSON(SvcData)

}
