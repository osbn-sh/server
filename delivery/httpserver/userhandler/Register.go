package userhandler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"ostadbun/param/userparam"
	"ostadbun/pkg/regexp"
	"ostadbun/pkg/richerror"

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

		var er richerror.RichError
		ew := errors.As(errR, &er)
		fmt.Println(ew, "this is oq3j")
		if ew {
			log.Println("root cause:", er.RootCause())

		}

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": errR.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
