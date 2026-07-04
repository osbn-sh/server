package userhandler

import (
	"fmt"
	"net/http"
	"ostadbun/param/userparam"
	"ostadbun/pkg/regexp"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) Register(c *fiber.Ctx) error {

	var user userparam.User

	err := c.BodyParser(&user)

	if err != nil {
		return fiber.ErrBadRequest
	}

	valid := regexp.IsEmailValid(user.Email)
	fmt.Println("validated")

	if !valid {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "email not valid",
		})
	}

	errR := h.userSvc.RegisterUserByEmailAndPassword(user)

	if errR != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": errR.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
