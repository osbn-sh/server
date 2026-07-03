package academic

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) UniversityGet(c *fiber.Ctx) error {

	target := c.Params("id")
	dta, err := h.academicService.UniversityGet(target)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}
	return c.JSON(dta)

}

func (h Handler) LessonGet(c *fiber.Ctx) error {

	target := c.Params("id")
	dta, err := h.academicService.LessonGet(target)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}
	return c.JSON(dta)

}

func (h Handler) MajorGet(c *fiber.Ctx) error {

	target := c.Params("id")
	dta, err := h.academicService.MajorGet(target)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}
	return c.JSON(dta)

}

func (h Handler) ProfessorGet(c *fiber.Ctx) error {

	target := c.Params("id")
	dta, err := h.academicService.ProfessorGet(target)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}
	return c.JSON(dta)

}
