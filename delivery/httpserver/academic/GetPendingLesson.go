package academic

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetPreRequites(c *fiber.Ctx) error {

	target, errINt := c.ParamsInt("id")

	if errINt != nil {
		return c.Status(http.StatusNotFound).SendString("id not number")
	}

	dta, err := h.academicService.GetPendingLessonPreReq(target)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}
	return c.JSON(dta)

}

func (h Handler) GetCoRequites(c *fiber.Ctx) error {

	target, errINt := c.ParamsInt("id")

	if errINt != nil {
		return c.Status(http.StatusNotFound).SendString("id not number")
	}

	dta, err := h.academicService.GetPendingLessonCoReq(target)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}
	return c.JSON(dta)

}
