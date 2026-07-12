package votehandler

import (
	"fmt"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) DeleteRate(c *fiber.Ctx) error {

	id, errID := c.ParamsInt("rate_id")

	if errID != nil {
		return richerror.Out(
			richerror.New("deleterate.delivery").WithMessage("id not found").WithKind(richerror.KindInvalid),
			c)
	}

	errSvc := h.voteSvc.RemoveRate(id)

	if errSvc != nil {
		return richerror.Out(
			errSvc,
			c)
	}

	fmt.Println(errSvc)
	return c.JSON(fiber.Map{
		"success": "true",
	})
}
