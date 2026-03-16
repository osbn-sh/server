package academic

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) Multi(c *fiber.Ctx) error {

	param := c.Params("param")
	id := c.Params("id")

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	logic := param == "university" || param == "lesson" || param == "major" || param == "professor"

	if !logic {
		return c.SendString("error")
	}

	f, g := h.academicService.MultiDepend(id_int, param)

	//return c.JSON(fiber.Map{
	//	"p": param,
	//	"i": id,
	//})
	if g != nil {
		return c.SendString("error")
	}
	return c.JSON(f)
}
