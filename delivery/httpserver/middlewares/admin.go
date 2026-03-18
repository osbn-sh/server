package middlewares

import (
	"fmt"
	"ostadbun/pkg/richerror"
	"ostadbun/service/userservice"

	"github.com/gofiber/fiber/v2"
)

func IsAdmin(usv userservice.User) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {

		usid := c.Locals("user_id")

		userID, can := usid.(string)

		if !can {
			if userID == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":     "Unauthorized",
					"row_error": can,
				})
			}
		}

		if userID == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":          "Unauthorized",
				"row_error_usid": userID,
			})
		}

		errDB := usv.AdminChecker(userID)

		if errDB != nil {
			fmt.Println("e", errDB.(richerror.RichError).Operation())
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":        "Unauthorized",
				"row_error_db": errDB.Error(),
			})
		}

		return c.Next()

	}
}
