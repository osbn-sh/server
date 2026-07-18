package votehandler

import (
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetOption(c *fiber.Ctx) error {

	SvcData, err := h.voteSvc.GetOptions()
	if err != nil {
		return richerror.Out(err, c)
	}

	return c.JSON(SvcData)

}
