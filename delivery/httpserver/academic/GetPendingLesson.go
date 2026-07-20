package academic

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetRequitesCheck(c *fiber.Ctx) error {

	dta, err := h.academicService.GetPendingLessonCheck()
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}
	return c.JSON(dta)

}
