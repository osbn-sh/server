package student

import (
	"ostadbun/pkg/httpstorage"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// PassLessonAdd @Summary		Register a passed lesson
// @Description	Register a lesson that the authenticated student has already passed.
// @Tags			Student
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		studentparam.StudentPassDetail	true	"Passed lesson information"
// @Success		200		{string}	string	"success"
// @Failure		400		{object}	richerror.UserReport
// @Failure		401		{object}	richerror.UserReport
// @Failure		500		{object}	richerror.UserReport
// @Router			/student/pass [post]
func (h Handler) PassLessonDelete(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, errID := strconv.Atoi(idStr)
	if errID != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is invalid")
	}

	userID, errN := httpstorage.Get(c, "user_id").Number()

	if errN != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user id not found",
		})
	}

	do, err := h.studentService.Remove(id, userID)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if do {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "حذف شد!",
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "برای این اطلاعات موردی یافت نشد!",
		})
	}
}
