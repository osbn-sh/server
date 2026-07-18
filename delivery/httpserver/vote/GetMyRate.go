package votehandler

import (
	"ostadbun/param/voteparam"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetMyRate(c *fiber.Ctx) error {

	user_id, errID := httpstorage.Get(c, "user_id").Number()
	if errID != nil {
		return richerror.Out(
			richerror.New("getmyrate.delivery").WithMessage("userID not found").WithKind(richerror.KindUnauthorized),
			c)
	}

	entity := c.Params("entity")

	if entity == "" {
		return richerror.Out(
			richerror.New("voteget.delivery").WithMessage("entity not found").WithKind(richerror.KindInvalid),
			c)
	}

	slug, errIDSlug := c.ParamsInt("rate")

	if errIDSlug != nil {
		return richerror.Out(
			richerror.New("voteget.delivery").WithMessage("rate not found").WithKind(richerror.KindInvalid),
			c)
	}

	data := voteparam.Vote{
		Target:   entity,
		TargetID: slug,
	}

	SvcData, err := h.voteSvc.GetMyRate(user_id, data)

	if err != nil {
		return richerror.Out(err, c)
	}

	return c.JSON(SvcData)

}
