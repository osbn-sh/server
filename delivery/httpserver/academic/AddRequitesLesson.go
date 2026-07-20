package academic

import (
	"net/http"
	"ostadbun/param/academicParam"
	"ostadbun/pkg/richerror"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) AddPreLessonRequites(c *fiber.Ctx) error {

	var data academicParam.AddRequiteLesson

	Err := c.BodyParser(&data)

	if Err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{})
	}

	err := h.academicService.AddPreRequitesLesson(data)

	if err != nil {
		return richerror.Out(err, c)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})

}

func (h Handler) AddCoLessonRequites(c *fiber.Ctx) error {

	var data academicParam.AddRequiteLesson

	Err := c.BodyParser(&data)

	if Err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{})
	}

	err := h.academicService.AddPreRequitesLesson(data)

	if err != nil {
		return richerror.Out(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})

}
