package manipulation

import (
	"github.com/gofiber/fiber/v2"
)

// BasicPermission godoc
// @Summary Checking permission
// @Description Check whether the current user has manipulation permission
// @Tags Permission
// @Produce json
// @Success 200 {object} error
// @Failure 401 {object} error
// @Router /manipulation/permission  [get]
func (h Handler) BasicPermission(c *fiber.Ctx) error {

	return c.SendString("you can manipulate ! ")

}
