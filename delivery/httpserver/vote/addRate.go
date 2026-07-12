package votehandler

import (
	"fmt"
	"ostadbun/param/voteparam"
	"ostadbun/pkg/httpstorage"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) AddRate(c *fiber.Ctx) error {

	userId, errUserId := httpstorage.Get(c, "user_id").Number()

	if errUserId != nil {
		return richerror.Out(
			richerror.New("addrate.delivery").WithMessage("user not found").WithKind(richerror.KindUnauthorized),
			c)
	}

	var data voteparam.Vote

	errBody := c.BodyParser(&data)

	if errBody != nil {
		return richerror.Out(
			richerror.New("addrate.delivery").WithMessage("error on parsing data").WithKind(richerror.KindInvalid),
			c)
	}

	fmt.Println(userId, data)
	errSvc := h.voteSvc.AddRate(userId, data)

	if errSvc != nil {
		return richerror.Out(
			errSvc,
			c)
	}

	return c.JSON(fiber.Map{
		"success": "true",
	})
}
