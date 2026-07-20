package academic

import (
	"net/http"
	"ostadbun/param/academicParam"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) LessonRequitesApproveOK(c *fiber.Ctx) error {

	var data academicParam.RequiteApprovement

	Err := c.BodyParser(&data)

	if Err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{})
	}
	var err error

	if data.CategorySTR == "pre" {
		err = h.academicService.ApprovementPreRequites(data, true)
	} else if data.CategorySTR == "co" {
		err = h.academicService.ApprovementCoRequites(data, true)
	} else {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})

}

func (h Handler) LessonRequitesApproveReject(c *fiber.Ctx) error {

	var data academicParam.RequiteApprovement

	Err := c.BodyParser(&data)

	if Err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{})
	}

	var err error

	if data.CategorySTR == "pre" {
		err = h.academicService.ApprovementPreRequites(data, false)
	} else if data.CategorySTR == "co" {
		err = h.academicService.ApprovementCoRequites(data, false)
	} else {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})

}
