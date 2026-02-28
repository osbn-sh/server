package academic

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) UniversityGet(c *fiber.Ctx) error {

	i := c.Params("id")

	id, err := strconv.Atoi(i)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	dta, err := h.academicService.UniversityGet(id)

	return c.JSON(dta)

}

func (h Handler) LessonGet(c *fiber.Ctx) error {

	i := c.Params("id")

	id, err := strconv.Atoi(i)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	dta, err := h.academicService.LessonGet(id)

	return c.JSON(dta)

}

func (h Handler) MajorGet(c *fiber.Ctx) error {

	i := c.Params("id")

	id, err := strconv.Atoi(i)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	dta, err := h.academicService.MajorGet(id)

	return c.JSON(dta)

}

func (h Handler) ProfessorGet(c *fiber.Ctx) error {

	i := c.Params("id")

	id, err := strconv.Atoi(i)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	dta, err := h.academicService.ProfessorGet(id)

	return c.JSON(dta)

}
