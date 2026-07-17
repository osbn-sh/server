package manipulation

import (
	"ostadbun/pkg/httpstorage"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetMyPending(c *fiber.Ctx) error {

	userId, err := httpstorage.Get(c, "user_id").Number()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	uni, errU := h.manipulSVC.GetPendingUniversity(userId)
	les, errL := h.manipulSVC.GetPendingLesson(userId)
	prof, errP := h.manipulSVC.GetPendingProfessor(userId)
	major, errM := h.manipulSVC.GetPendingMajor(userId)

	if errM != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errM.Error(),
			"code":  "major",
		})
	}

	if errL != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errL,
			"code":  "lesson",
		})
	}

	if errU != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errU,
			"code":  "university",
		})
	}

	if errP != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errP,
			"code":  "prof",
		})
	}

	if errU != nil && errL != nil && errP != nil && errM != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Something went wrong",
			"errorsOn": map[string]string{
				"University": errU.Error(),
				"Professor":  errP.Error(),
				"Lesson":     errL.Error(),
				"Major":      errM.Error(),
			},
		})
	}

	Data := IProvide{
		Lesson:     les,
		University: uni,
		Professor:  prof,
		Major:      major,
	}

	return c.JSON(Data)

}
